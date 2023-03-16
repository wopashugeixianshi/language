package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//获取当前目录下的所有文件或目录信息
	filepath.Walk("/proc/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		fmt.Println(path)        //打印path信息
		fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})
}
