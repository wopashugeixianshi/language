package main

import (
	"fmt"
	"time"
)

func main() {
	/* 100 0000 --- 80ms*/
	start := time.Now()
	fmt.Println("start, accumulate cost", time.Since(start))
	for i := 0; i < 10000000; i++ {
		fmt.Sprint("%d", i)
	}
	fmt.Println("  end, accumulate cost", time.Since(start))
}
