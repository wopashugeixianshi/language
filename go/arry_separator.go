package main

import (
	"fmt"
	"strings"
)

func main() {
	var src = []int{30024, 200, 300}
	var temp = make([]string, len(src))
	var sort = make(map[int]bool)
	var k int
	for _, v := range src {
		if v == 200 {
			sort[v] = true
    		temp[k] = fmt.Sprintf("%d", v)
			k++
		}
	}
	fmt.Println(len(sort))
	var result = strings.Join(temp, ",")
	fmt.Println(result)
}
