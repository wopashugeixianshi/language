package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ncw/directio"
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
	in, err := directio.OpenFile(pathname, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("[%s]\n", pathname)
	block := directio.AlignedBlock(directio.BlockSize)
	responseByte, err := io.ReadFull(in, block)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected EOF") && responseByte != 0 {
			fmt.Println(err, responseByte)
			in.Close()
			return nil
		}
		in.Close()
		return err
	}
	in.Close()
	return nil
}

func main() {
	GetAllFile("/host/proc/")
}
