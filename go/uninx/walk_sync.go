package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
	"strings"
)

func Fsync(f *os.File) error {
	if err := unix.Fsync(int(f.Fd())); err != nil {
		return err
	}
	return nil
}

func fadviseLib(f *os.File, fsync bool, stat os.FileInfo) (err error) {
	if fsync {
		if stat.Mode().Perm()&os.FileMode(0200) == os.FileMode(0200) { // 验证是否有可写权限
			if err = Fsync(f); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	var fileCount int
	var dirCount int
	var totalCount int
	//获取当前目录下的所有文件或目录信息
	filepath.Walk("/proc",func(path string, info os.FileInfo, err error) error{
		fmt.Println(path) //打印path信息
		//fmt.Println(info.Name()) //打印文件或目录名
		totalCount++
		if info != nil && !info.IsDir() {
			fileCount++
			if strings.Contains(path, "/fd/8") {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				return nil
			}
			err = fadviseLib(f, true, info)
			if err != nil {
				fmt.Println(path, err)
				return nil
			}
			f.Close()
		}
		if info != nil && info.IsDir() {
			dirCount++
		}
		return nil
	})
	fmt.Println("file :", fileCount, " dir :", dirCount, " total :", totalCount)
}
