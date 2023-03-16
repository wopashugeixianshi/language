package main

import (
	"fmt"
)

func main() {
	var tmp [4]byte
	tmp[0] = 0
	tmp[1] = 1
	tmp[2] = 2
	tmp[3] = 3

	fmt.Println("tmp:", tmp)
	fmt.Println("tmp[:]:", tmp[:])
	fmt.Println("tmp[0:]:", tmp[0:])
	fmt.Println("tmp[1:]:", tmp[1:])
	fmt.Println("tmp[:3]:", tmp[:3])
	fmt.Println("tmp[:4]:", tmp[:4])
	fmt.Println("tmp[1:2]:", tmp[1:2])
}
