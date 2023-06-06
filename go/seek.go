package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fd, err := os.Open("./test")
	if err != nil {

	}
	offset, _ := fd.Seek(0, io.SeekStart)
	fmt.Println("start_offset:", offset)
	offset, _ = fd.Seek(0, io.SeekCurrent)
	fmt.Println("cur_offset:", offset)
	offset, _ = fd.Seek(0, io.SeekEnd)
	fmt.Println("end_offset:", offset)

	offset, err = fd.Seek(0, io.SeekStart)
	fmt.Println("start_offset:", offset, err)
	content := make([]byte, 10)
	n, err := fd.Read(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("n:", n)
	fd.Close()
	if fd != nil {
		fmt.Println("fd != nil")
	}
	offset, err = fd.Seek(0, io.SeekCurrent)
	fmt.Println("cur_offset:", offset, err)
	n, err = fd.Read(content)
	fmt.Println(n, err)
}
