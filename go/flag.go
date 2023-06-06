package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var delay time.Duration
	var delay1 int64
	//flag.DurationVar(&delay, "container-event-delay", 4*time.Second, "container event delay")
	flag.DurationVar(&delay, "container-event-delay", 4, "container event delay")
	flag.Int64Var(&delay1, "container-event-delay1", 60, "container event delay1")
	flag.Parse()

	fmt.Println(delay*time.Minute, delay1)
	fmt.Println("-------visit----------")
	flag.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value)
	})
	fmt.Println("-------visit_all---------")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value)
	})
	fmt.Println("---------------------")
	fmt.Println(delay, int64(delay))
	second := int64(delay/time.Second)
	fmt.Println(second)
	//if int64(delay) > 5000000000 {
	if int64(delay) > 5000000000 {
		delay = time.Duration(5000000000)
	}
	timer := time.NewTicker(delay)
	select {
	case <-timer.C:
		fmt.Println("timeout")
	}
	fmt.Printf("Container event check time %v, delayed update time %vs\n", delay, delay1)
	fmt.Println(delay, int64(delay))
	fmt.Println(delay1)
}
