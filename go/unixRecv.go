package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
)


const recvServer string = "/tmp/recvServer.sock"

func main() {
	pprofStart()
	listenDP()
}

func pprofStart() {
	go func (){
		http.ListenAndServe("0.0.0.0:9090", nil)
	}()
}

func listenDP() {
	os.Remove(recvServer)
	var conn *net.UnixConn
	kind := "unixgram"
	addr := net.UnixAddr{recvServer, kind}
	defer os.Remove(recvServer)
	conn, _ = net.ListenUnixgram(kind, &addr)
	defer conn.Close()

	for {
		var buf [2048]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println("Read message err:", err)
		} else {
			log.Println("Recv bytes:", n)
			log.Println("Recv msg:", buf)
		}
	}
}
