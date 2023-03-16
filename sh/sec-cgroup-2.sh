#!/bin/bash -e
usage() {
  script=./$(basename $0)
  cat <<-EOF
	$(tput bold)用法:$(tput sgr0)
	  # 采集agent内存，默认24小时后自动结束
	  $(tput setaf 2)${script} $(tput sgr0)

	  # 采集server内存
	  $(tput setaf 2)${script} server$(tput sgr0)

	  # 采集agent内存，不会自动结束
	  $(tput setaf 2)${script} infinity$(tput sgr0)

	  # 采集server内存，不会自动结束
	  $(tput setaf 2)${script} server infinity$(tput sgr0)

	  # 采集agent内存详情
	  $(tput setaf 2)${script} stat$(tput sgr0)
EOF
}

proc=sec_hades
infinity=false
dura=24h
printMemStat=false

hms(){
    x=${1//Y/*31536000+}
    x=${x//M/*2628000+}
    x=${x//D/*86400+}
    x=${x//h/*3600+}
    x=${x//m/*60+}
    x=${x//s/+}
    x=${x%+}
    echo "$((x))"
}

printFormat(){
  printf "%-7s : %s\n" "$1" "$2"
}

while (( "$#" )); do
  case "$1" in
    server)
      proc=sec_server
      ;;
    scanner)
      proc=sec_scanner
      ;;
    app)
      proc=sec_app
      ;;
    -s)
      proc=$2
      shift
      ;;
    -d)
      dura=$2
      shift
      ;;
    infinity)
      infinity=true
      ;;
    stat)
      printMemStat=true
      ;;
    -h|--help)
      usage
      exit
      ;;
    *)
      echo "args format error $1"
      usage
      exit
  esac
  shift
done

procPs=$(ps aux | grep ${proc} | grep -v grep|grep -v "\-s ${proc}")
procPid=$(echo ${procPs}|tr -s ' '|cut -d " " -f2)

if [ -z "${procPid}" ]; then
  echo "proc:"$proc
  echo -ne "\033[31m"
  echo "can't find process ${proc}"
  echo -ne "\033[0m\n"
  exit
fi

printFormat "pid" $procPid
printFormat "proc" $procPs

pidCgroupPath=/proc/${procPid}/cgroup
podCgroupPath=$(awk -F ":" '/:memory:/{print $3}' <${pidCgroupPath})
podCgroupPath="/proc/1/root/sys/fs/cgroup/memory$podCgroupPath"
if [ ! -d "${podCgroupPath}" ]; then
  echo -e "\n""container not exits"
  exit
fi
printFormat "cgroup" ${podCgroupPath}
cd "${podCgroupPath}"

if [ ! -f "memory.stat" ]; then
  echo -e "\n""container not exits"
  exit
fi

if $printMemStat; then
  echo "path    : "${podCgroupPath}
  grep "total" memory.stat | grep -vE "pg| 0" |
    awk '{ split( "B KB MB GB" , v ); s=1; while( $2>1024 ){ $2/=1024; s++ } printf $1" " "%.2f%s  \n",$2, v[s]}'

  pageSize=$(getconf PAGESIZE)
  grep "total" memory.stat | grep pg | grep -v " 0" |
    awk -v pageSize=$pageSize '{ split( "B KB MB GB" , v ); s=1;$2=pageSize * $2; while( $2>1024 ){ $2/=1024; s++ } printf $1" " "%.2f%s  \n",$2, v[s]}'
  exit
fi

if ! $infinity; then
  nowTimestamp=$(date +%s)
  endTimestamp=$((nowTimestamp + $(hms $dura)))
  printFormat "start" "$(date +"%m/%d %H:%M:%S" -d "@$nowTimestamp")"
  printFormat "end" "$(date +"%m/%d %H:%M:%S" -d "@$endTimestamp")"
fi

while true; do
  if [ ! -f "memory.stat" ]; then
    echo -e "\n""container not exits"
    exit
  fi

  echo -en "\n""$(date +"%m/%d %H:%M:%S")""  "
  grep active memory.stat | grep -v " 0" | grep total |
    awk '{ split( "B KB MB GB" , v ); s=1; while( $2>1024 ){ $2/=1024; s++ } printf $1" " "%.2f%s  ",$2, v[s]}'
  awk '{ split( "B KB MB GB" , v ); s=1; while( $1>1024 ){ $1/=1024; s++ } printf "usage_in_bytes %.2f%s  ",$1, v[s] }' memory.usage_in_bytes
  file="memory.kmem.usage_in_bytes" && [[ -f $file ]] && awk '{ split( "B KB MB GB" , v ); s=1; while( $1>1024 ){ $1/=1024; s++ } printf "kmem %.2f%s  ",$1, v[s] }' $file

  if  ! $infinity && [ "$(date +%s)" -gt "${endTimestamp}" ]; then
    echo -e "\n到达结束时间"
    exit 0
  fi
  sleep 5
done

# 后台常驻启动方式
#nohup sh sec-monitor-cgroup.sh > ./sec-monitor-agent.log 2>&1 &
#nohup sh sec-monitor-cgroup.sh server > ./sec-monitor-server.log 2>&1 &
#nohup sh sec-monitor-cgroup.sh infinity > ./sec-monitor-agent.log 2>&1 &
