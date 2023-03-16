package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const ctrlServer string = "/tmp/ctrl_listen.sock"
const ctrlClient string = "/tmp/ctrl_send.sock"

func connectDP() {
	os.Remove(ctrlClient)
	var conn *net.UnixConn
	var err error
	kind := "unixgram"
	laddr := net.UnixAddr{ctrlClient, kind}
	raddr := net.UnixAddr{ctrlServer, kind}
	conn, err = net.DialUnix(kind, &laddr, &raddr)
	if err != nil {
		fmt.Println("conn:", err)
	}
	defer os.Remove(ctrlClient)
	defer conn.Close()

	for {
		var buf [4096]byte
		buf[0] = 0x11
		buf[1] = 0x12
		n, err := conn.Write(buf[:])
		if err != nil {
			fmt.Printf("Read message err: %v\n", err)
		} else {
			fmt.Println("snd num:", n);
		}
		
		time.Sleep(time.Duration(2)*time.Second)
	}
}

func main() {
	connectDP();
}
