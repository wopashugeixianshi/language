package main

import (
	"fmt"
)

type namespace struct {
	hostname string
	pid      int
}
/* 指针(*)优先级小于后缀运算符(->) */
func main() {
	namespace := &namespace{
		hostname: "ll-master",
		pid:      100,
	}
	tmp := namespace
	fmt.Println((*tmp).hostname)

	var temp interface{}
	temp = "haha"
	fmt.Println(temp.(string))
}
