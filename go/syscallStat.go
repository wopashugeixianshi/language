package main

import (
	"fmt"
	"syscall"
)

func main() {
	var err error
	var HostPidNS uint32
	var HostNetNS uint32
	for i := 1; i < 10; i++ {
		HostPidNS, err = getNS(fmt.Sprintf("/proc/%d/%s", i, "ns/pid"))
		if err != nil {
			fmt.Println(err)
			continue
		}
		HostNetNS, err = getNS(fmt.Sprintf("/proc/%d/%s", i, "ns/net"))
		if err != nil {
			fmt.Println(err)
		}
		
		fmt.Printf("i = %d, HostPidNS = %d, HostNetNS = %d\n", i, HostPidNS, HostNetNS)
	}
}

func getNS(path string) (uint32, error) {
	var statT syscall.Stat_t
	if err := syscall.Stat(path, &statT); err != nil {
		return 0, err
	}
	return uint32(statT.Ino), nil
}
