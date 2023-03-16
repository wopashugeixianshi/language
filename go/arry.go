package main

import (
	"fmt"
	"unsafe"
)

type PreProcessContainer struct {
	EventType     int 
	Action        int 
	CycleTime	  int64
}

var intArry []int
var tmp map[string]PreProcessContainer

func main() {
	//intArry[0] = 100
//	fmt.Println(intArry)
	tmp = make(map[string]PreProcessContainer, 1000)
	fmt.Println(unsafe.Sizeof(tmp))
	tmp["haha"] = PreProcessContainer {
		EventType: 100,
		Action: 1,
		CycleTime : 1000,
	}
	value, ok := tmp["hehe"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("none exist")
	}
	fmt.Println(tmp)
}
