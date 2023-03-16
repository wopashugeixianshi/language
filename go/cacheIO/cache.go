package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+fi.Name())
			GetAllFile(pathname + fi.Name() + "/")
		} else {
			//fmt.Printf("[%s]\n", pathname+fi.Name())
			ReadFileinfo(pathname + fi.Name())
		}
	}
	return err
}

func ReadFileinfo(pathname string) error {
	in := strings.NewReader(pathname)
	fmt.Printf("[%s]\n", pathname)
	block := make([]byte, 4096)
	responseByte, err := io.ReadFull(in, block)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") && responseByte != 0 {
			//fmt.Println(err, responseByte)
			return nil
		}
		return err
	}
	return nil
}

func main() {
	GetAllFile("/host/proc/")
}
