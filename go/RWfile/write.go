package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//file, err := os.Open("./file")
	file, err := os.OpenFile("./file", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		writeFile(file)
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond)
		//time.Sleep(time.Microsecond)
	}
	file.Close()
}

func writeFile(file *os.File) {
	if file == nil {
		fmt.Println("file == nil")
	}
	num, err := file.WriteString("hello,world\n")
	if err != nil {
		fmt.Println(num, err)
	}
}
