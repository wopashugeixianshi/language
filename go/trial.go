package main

import (
	"fmt"
)

func main() {
	var vari = "haha"
	loop(vari)
}

func loop(args ...interface{}) {
	fmt.Println(args)
	fmt.Println(args[0].(string))

	ddd := fmt.Sprintf("%s", args[0])
	fmt.Println("printlnddd:", ddd)
	fmt.Println("println:", []byte(ddd))
	fmt.Printf("%%v:%v\n", []byte(ddd))
	fmt.Printf("%%T:%T\n", []byte(ddd))
}
