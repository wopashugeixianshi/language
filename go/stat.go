package main

import (
	"fmt"
	"os"
	"path/filepath"
	//"strconv"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 8000; i++ {
		files, _ := filepath.Glob(filepath.Join("", "/proc/[0-9]*"))
		for _, file := range files {
			data, _ := os.ReadFile(filepath.Join(file, "stat"))
			info := strings.FieldsFunc(string(data), func(r rune) bool {
				return r == '(' || r == ')'
			})
			if len(info) != 3 {
				fmt.Println("len:", len(info))
				continue
			}

			for i, v := range info {
				info[i] = strings.TrimSpace(v)
			}
			//fmt.Println(info[2])
			//fieldSlice := strings.Split(info[2], " ")
			//strconv.Atoi(fieldSlice[1])
			//strconv.Atoi(fieldSlice[4])

			var state byte
			var grp int
			var session int
			var gid int
			var flags uint32
			var minFlt uint64
			var cminFlt uint64
			var majFlt uint64
			var cmajFlt uint64
			var uTime int64
			var sTime int64
			var PPID int
			var TTY int
			fmt.Sscanf(
				info[2], "%c %d %d %d %d %d %d %d %d %d %d %d %d",
				&state,
				&PPID, &grp, &session,
				&TTY, gid, flags,
				minFlt, cminFlt, majFlt, cmajFlt,
				uTime, sTime,
			)
			//fmt.Println(PPID, TTY)
		}
	}
	for {
	}
	time.Sleep(time.Second * 10)
}
