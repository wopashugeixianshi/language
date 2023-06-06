package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"golang.org/x/sys/unix"
)

func main() {
	//defer func() {
	//	recover()
	//}()
	for {
		file, err := os.Open("../file")
		if err != nil {
			panic(err)
		}
		buffer := make([]byte, 12)
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				panic(err)
			}
		}
		if n == 13 {
			fmt.Printf("Read from file: %s", buffer[:n])
		}
		unix.Fsync(int(file.Fd()))
		file.Close()
		time.Sleep(time.Second)
	}
}
