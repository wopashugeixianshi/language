package main

import (
	"fmt"
	"os"
//	"strings"
)


func main() {
	os.Setenv("GODEBUG", "cgocheck=0")
	fmt.Println(os.Getenv("GODEBUG"))
	for _, e := range os.Environ() {
		//pair := strings.Split(e, "=")
		//fmt.Println(pair[0])
		fmt.Println(e)
	}
}
