package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(getDeviceNameNotVirtual())
}

func getDeviceNameNotVirtual() (name string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return name, fmt.Errorf("get interfaces data err:", err)
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		//return iface.Name, err
		fmt.Println(iface.Name)
	}
	return name, err
}
