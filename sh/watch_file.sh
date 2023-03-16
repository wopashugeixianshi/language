#####################
## 脚本说明：监控目录文件大小，当超过设置大小时，清空文件.
## 容量限制 ：请调整 cap  变量中的数值，单位为K
## 监控路径 ：请调整 arrlogFile 变量中的值
#####################
#======================定义全局变量 begin===========================
#限制的容量大小单位k，配置100M
cap=102400
#监控路径,多个路径以空格隔开
#arrlogFile=("/dosec/log/waf/dp.log" )
arrlogFile=("/home/dosec/language/sh/log/" )
#host1="192.168.4.81"
host1="192.168.138.129"
hosts=($host1)
#======================定义全局变量 end=============================
#循环迭代配置的扫描目录 启动多个线程扫描任务
for ((j=0;j<${#hosts[@]};j++))
do
{
for ((i=0;i<${#arrlogFile[@]};i++)) do
 #迭代当前host
 currentHost=${hosts[$j]}
 #迭代当前目录文件
 currentFile=${arrlogFile[$i]}
 echo $currentHost---$currentFile
 if ssh $currentHost test -e $currentFile; then #判断在当前host上 目录是否存在 如果存在执行操作
 curFileNum=`ssh $currentHost "cd $currentFile;ls -l|wc -l"` 
 #当前的文件夹容量大小
 curcap=`ssh $currentHost "cd $currentFile;du -sk|awk '{print $1}'" | awk '{print $1}'` 
 #平均每个文件大小，如果超过平均大小则清空文件
 echo $curcap
 fileAvg=$((cap/(curFileNum-1)))
 echo 'avg file size--->'$fileAvg
 #暂存遍历的单个文件大小  
 size=0  
 echo fileAvg$fileAvg
 if [[ $curcap -gt $cap ]] ; then  #当监控目录文件大小大于设置的控制文件大小时去处理文件.
      echo '清理开始'
      for file in `ssh $currentHost "cd $currentFile;ls -l|awk '{print $9}'" | awk '{print $9}'`  #列出目录下所有文件
        do  
                if ssh $currentHost test -f $currentFile/$file;  then
                        echo 'current file->'$file 
                        size=`ssh $currentHost "cd $currentFile;du -sk $file | awk '{print $1}'" | awk '{print $1}'`  #查看文件大小
                        if [[ $size -gt $fileAvg ]] ; then  #当文件大小大于控制文件在这个目录的平均文件大小时清理文件
                             ssh $currentHost "cd $currentFile;cat /dev/null > $file" #清空文件
                        fi  
                fi  
     done   
   fi
 fi
 done
} &
done
wait
echo "监控结束"
#直接通过命令 
#chmod 750 monitorLogs.sh
#./monitorLogs.sh
#或sh monitorLogs.sh 
#执行
