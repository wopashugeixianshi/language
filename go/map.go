package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type PreProcessContainer struct {
	EventType int
	Action int
	CycleTime int64
	Event map[string]interface{}
}

var PreProcessContainers sync.Map

func main() {
	//var event map[string]int
	var event1 map[string]int
	event1 = make(map[string]int, 10)
	event1["haha"] = 100
	event1["hehe"] = 200
	event1["xixi"] = 300
	//event = make(map[string]int,1)
	event := event1

	//PreProcessContainers.Store("haha", "hehe")
	//PreProcessContainers.Store("heihei", "hehe")
	//PreProcessContainers.Store("xixi", "hehe")
	//fmt.Println(PreProcessContainers)
	PreProcessContainers.Range(func(k, v interface{}) bool {
		if k.(string) == "heihei" {
			v = "wotule"
			return false
		}
		fmt.Println("haha")
		return true
	})

	PreProcessContainers.Range(func(k, v interface{}) bool {
		fmt.Println(k.(string), v.(string))
		return true
	})

	pre := PreProcessContainer{}
	fmt.Println("struct sizeof :", unsafe.Sizeof(pre))
	fmt.Println("map sizeof :", unsafe.Sizeof(PreProcessContainers))
	fmt.Println(event)
}
