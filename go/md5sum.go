package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

func printMemInfo(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}

func getFileMD5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	printMemInfo("开始")

	tmp, _ := getFileMD5("./data")
	fmt.Println(tmp)
	printMemInfo("结束")
}
