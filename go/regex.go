package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "prOtocol 1"
	regex := "(?i)Protocol[\\s\\S]*(2|1)"
	reg, err:= regexp.Compile(regex)
	if err != nil {
		fmt.Println(err)
	}

	match := reg.FindStringSubmatch(str1)
	fmt.Println(len(match), "\t", match)
	for _, v := range match {
		fmt.Println(v)
	}

	fmt.Println("--------------------------------------------------")
	str := "/home/dosec/language/go" + "/[^/]+" + "/[^/]+" + "/test"
	//fmt.Println(str)
	reg = regexp.MustCompile(str)
	str1 = "/home/dosec/language/go/1/2/test"
	match = reg.FindStringSubmatch(str1)
	for i, v := range match {
		fmt.Println(i, v)
	}
}
