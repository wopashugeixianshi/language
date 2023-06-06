package main

import (
	"fmt"
	"syscall"
)

func main() {
	path := "/proc/2019/ns/pid"
	var statT syscall.Stat_t
	if err := syscall.Stat(path, &statT); err != nil {
		fmt.Println("err :", err)
		return 
	}

	fmt.Println(uint32(statT.Ino))
}
