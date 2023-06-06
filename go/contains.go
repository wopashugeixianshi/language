package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "bclinux8"
	var str2 = "bclinux8"
	var str3 = "bclinux7"
	var str4 = "bclinux8."
	var str5 = "bclinux8.1"

	if strings.Contains(str1, str2) {
		fmt.Println("true")
	}
	fmt.Println("-------------------------")
	if strings.Contains(str3, str2) {
		fmt.Println("true")
	}
	fmt.Println("-------------------------")
	if strings.Contains(str4, str2) {
		fmt.Println("true")
	}
	fmt.Println("-------------------------")
	if strings.Contains(str5, str2) {
		fmt.Println("true")
	}
}
