package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "请输入 -cpuprofile 指定性能分析文件名称")
	var memprofile = flag.String("memprofile", "", "请输入 -memprofile 指定内存分析文件名称")
	//在所有flag都注册之后,调用flag.Parse()
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal("could not create cpuprofile:", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile:", err)
	}
	defer pprof.StopCPUProfile()

	f, err = os.Create(*memprofile)
	if err != nil {
		log.Fatal("could not create memprofile:", err)
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write MEM profile:", err)
	}
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Printf("sum=%d\n", sum)
}
