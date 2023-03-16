package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//exePath := filepath.Join("/proc/1803", "exe")
	//fmt.Println(exePath)
	//exe, err := os.Readlink(exePath)
	//if err != nil {
	//
	//}
	if exe, err := os.Readlink(filepath.Join("/proc/1804", "exe")); err != nil {
		fmt.Println("haha:", err)
	} else {
		fmt.Println(exe)
	}
}
