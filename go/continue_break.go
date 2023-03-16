package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		//label2: 0 1 3 4    0 1 3 4
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue 
			}
			fmt.Println("j = ", j)
		}
	}
}
