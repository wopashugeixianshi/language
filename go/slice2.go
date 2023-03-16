package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//b := make([]byte, 0, 192)
	b := make([]byte, 192)
	//println(len(b), cap(b))
	r, err := os.Open("/proc/1931/status")
	if err != nil {
		return
	}
	// 读取文件大小
	//println(b)
	println("len :", len(b[len(b):cap(b)]))
	fmt.Printf("%v\n", b[len(b):cap(b)])
	//n, err := r.Read(b[len(b):cap(b)])
	n, err := r.Read(b)
	println(len(b), cap(b), n)
	fmt.Printf("%v\n", b[len(b):cap(b)])
	fmt.Printf("%v\n", b[:cap(b)])
	fmt.Printf("%v\n", b[:])
	
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			return
		}
	}
	//b = b[:len(b)+n]
	println(len(b), cap(b), n)
	println(b[:])
	fmt.Printf("%v", b[:])
	r.Close()
}

