package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func workOnce(wg *sync.WaitGroup) {
	counter()
	wg.Done()
}

func main() {
	/* 返回值类型是*string */
	var cpuProfile = flag.String("cpuprofile", "", "请输入 -cpuprofile 指定CPU分析文件名")
	var memProfile = flag.String("memprofile", "", "请输入 -memprofile 指定MEM分析文件名")
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}

	wg.Wait()
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
