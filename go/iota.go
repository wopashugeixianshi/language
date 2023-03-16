package main

import(
	"fmt"
)

type EventType int
const  (
	_ EventType = iota
	ProcessCreationEvent 
	NetCreationEvent
	ProcessExitEvent
)

func main() {
	fmt.Println(ProcessExitEvent, NetCreationEvent, ProcessCreationEvent)
}
