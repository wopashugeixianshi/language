package main

//#include "/home/dosec/Go/src/neuvector/defs.h"
import "C"

import (
	"encoding/binary"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"unsafe"
)

const ctrlServer string = "/tmp/ctrl_listen.sock"

func ParseDPMsgHeader(msg []byte) *C.DPMsgHdr {
	var hdr C.DPMsgHdr
	//log.Info("ParseDPMsgHeader:", string(msg))
	hdrLen := int(unsafe.Sizeof(hdr))
	if len(msg) < hdrLen {
		log.Println("Short header len:", len(msg))
		return nil
	}

	r := bytes.NewReader(msg)
	binary.Read(r, binary.BigEndian, &hdr)
	if int(hdr.Length) != len(msg) {
		log.Println("Wrong message length.", "kind:", hdr.Kind, "expect:", hdr.Length, "actual:", len(msg))
		return nil
	}

	return &hdr
}

func listenDP() {
	os.Remove(ctrlServer)
	var conn *net.UnixConn
	var err error
	kind := "unixgram"
	addr := net.UnixAddr{ctrlServer, kind}
	defer os.Remove(ctrlServer)
	conn, err = net.ListenUnixgram(kind, &addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		var buf [4096]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("Read message err: %v\n", err)
		} else {
			hdr := ParseDPMsgHeader(buf[:n])
			if hdr == nil {
				return
			}
				switch int(hdr.Kind) {
				case C.DP_KIND_APP_UPDATE:
					log.Println("hdr.Kind = ", C.DP_KIND_APP_UPDATE, " num = ", n)
				case C.DP_KIND_THREAT_LOG:
					log.Println("hdr.Kind = ", C.DP_KIND_THREAT_LOG, " num = ", n)
				case C.DP_KIND_CONNECTION:
					log.Println("hdr.Kind = ", C.DP_KIND_CONNECTION, " num = ", n)
				case C.DP_KIND_FQDN_UPDATE:
					log.Println("hdr.Kind = ", C.DP_KIND_FQDN_UPDATE, " num = ", n)
				}
		}
	}
}

func main() {

	listenDP();
}
