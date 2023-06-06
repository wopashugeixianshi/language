package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

var (
	device       string = "ens33"
	snapshot_len int32  = 1024
	promiscuous  bool   = true
	err          error
	timeout      time.Duration = 1 * time.Second
	handle       *pcap.Handle
)

func main() {
	// 打开某一网络设备
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	//defer handle.Close()
	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	filterStr := "not host 192.168.138.129"
	if err := handle.SetBPFFilter(filterStr); err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		time.Sleep(1000*time.Second)
		handle.Close()
	}()
	fmt.Println("starting capture packet")
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
	fmt.Println("end")
}
