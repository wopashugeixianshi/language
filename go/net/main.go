package main

import (
	"fmt"
	"syscall"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

type Addr struct {
	IP   string
	Port uint32
}

type connTmp struct {
	fd       int
	family   int
	sockType uint32
	laddr    Addr
	raddr    Addr
	status   string
	pid      int
	inode    uint32
}

var TCPIntStatuses = map[int]string{
	//netlink.TCP_ESTABLISHED: "ESTABLISHED",
	//netlink.TCP_SYN_SENT:    "SYN_SENT",
	//netlink.TCP_SYN_RECV:    "SYN_RECV",
	//netlink.TCP_FIN_WAIT1:   "FIN_WAIT1",
	//netlink.TCP_FIN_WAIT2:   "FIN_WAIT2",
	//netlink.TCP_TIME_WAIT:   "TIME_WAIT",
	//netlink.TCP_CLOSE:       "CLOSE",
	//netlink.TCP_CLOSE_WAIT:  "CLOSE_WAIT",
	//netlink.TCP_LAST_ACK:    "LAST_ACK",
	//netlink.TCP_LISTEN:      "LISTEN",
	//netlink.TCP_NEW_SYN_REC: "CLOSING",

	1:  "ESTABLISHED",
	2:  "SYN_SENT",
	3:  "SYN_RECV",
	4:  "FIN_WAIT1",
	5:  "FIN_WAIT2",
	6:  "TIME_WAIT",
	7:  "CLOSE",
	8:  "CLOSE_WAIT",
	9:  "LAST_ACK",
	10: "LISTEN",
	11: "CLOSING",
}

func main() {
	conn, err := processInetDiag()
	if err != nil {
		return
	}

	for _, v := range conn {
		fmt.Println(v)
	}
}

func processInetDiag() ([]connTmp, error) {
	var ret []connTmp

	protos := []int{unix.IPPROTO_UDP, unix.IPPROTO_TCP}
	for _, proto := range protos {
		//socketDiags, err := netlink.SocketDiagTCP(syscall.AF_INET, uint8(proto))
		socketDiags, err := netlink.SocketDiagTCP(syscall.AF_INET)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(len(socketDiags))
		for _, socketDiagInfo := range socketDiags {
			la := Addr{
				IP:   socketDiagInfo.ID.Source.String(),
				Port: uint32(socketDiagInfo.ID.SourcePort),
			}

			ra := Addr{
				IP:   socketDiagInfo.ID.Destination.String(),
				Port: uint32(socketDiagInfo.ID.DestinationPort),
			}

			state := int(socketDiagInfo.State)
			status := "NONE"
			sockType := syscall.AF_UNSPEC

			if proto == unix.IPPROTO_TCP {
				status = TCPIntStatuses[state]
				sockType = syscall.SOCK_STREAM
			}
			//else if proto == unix.IPPROTO_UDP {
			//	sockType = syscall.SOCK_DGRAM
			//	if _, ok := UDPIntStatuses[state]; ok {
			//		if _, ok := UDPIntStatuses[state]; ok {
			//			status = UDPIntStatuses[state]
			//		}
			//	}

			//}
			ret = append(ret, connTmp{
				fd:       0,
				family:   syscall.AF_INET,
				sockType: uint32(sockType),
				laddr:    la,
				raddr:    ra,
				status:   status,
				pid:      -1,
				inode:    socketDiagInfo.INode,
			})
		}
	}
	return ret, nil
}
