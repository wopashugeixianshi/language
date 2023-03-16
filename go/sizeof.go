package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type test struct {
	isServ bool
	sync.Mutex
}

type test1 struct {
	isDeleted bool
	isHandle bool
	test
	startTime int
}

type PreProcessor struct{
	eventType	  int
	action        int
	cycleTime	  int64
	containerID   string
}

func main() {
	tmp := test{}
	tmp1 := test1{}
	tmp2 := PreProcessor{}
	var intVar int
	var int64Var int64
	var boolVar bool
	var strVar string
	fmt.Println(unsafe.Sizeof(tmp))
	fmt.Println(unsafe.Sizeof(tmp1))
	fmt.Println(unsafe.Sizeof(intVar))
	fmt.Println(unsafe.Sizeof(int64Var))
	fmt.Println(unsafe.Sizeof(boolVar))
	fmt.Println("string :", unsafe.Sizeof(strVar))
	fmt.Println(unsafe.Sizeof(tmp2))
}
