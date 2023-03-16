package main

import (	
	"fmt"
)

type test struct {
	id int
}

type test1 struct {

}

func main() {
	var x interface{}
	x = test{
		id: 100,
	}
	_, ok := x.(string) 
	if ok {
		fmt.Println("ok")
	}
	_, ok = x.(test1)
	if ok {
		fmt.Println("ok")
	}
}
