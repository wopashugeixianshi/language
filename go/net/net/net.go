package net

import (
	"sync"
	"syscall"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)


type Addr struct {
	IP string
	Port  uint32
}

type connTmp struct {
	fd int
	family int
	sockType uint32
	laddr Addr
	laddr Addr
	status stringa
	pid int
	inode uint32
}

type inodeMap struct {
	pid int
	ns string
}

var TCPIntStatuses = map[int]string{
	netlink.TCP_ESTABLISHED: "ESTABLISHED",
	netlink.TCP_SYN_SENT:    "SYN_SENT",
	netlink.TCP_SYN_RECV:    "SYN_RECV",
	netlink.TCP_FIN_WAIT1:   "FIN_WAIT1",
	netlink.TCP_FIN_WAIT2:   "FIN_WAIT2",
	netlink.TCP_TIME_WAIT:   "TIME_WAIT",
	netlink.TCP_CLOSE:       "CLOSE",
	netlink.TCP_CLOSE_WAIT:  "CLOSE_WAIT",
	netlink.TCP_LAST_ACK:    "LAST_ACK",
	netlink.TCP_LISTEN:      "LISTEN",
	netlink.TCP_NEW_SYN_REC: "CLOSING",
}

var UDPStatuses = map[string]string{
	"01": "ESTABLISHED",
	"07": "LISTEN",
}

//"&{2 1 2 0 {21186 8000 10.208.136.12 100.69.239.4 0 [868200256 6]} 4313 0 0 0 2835862684}"
//只支持inet family
//func getDockerSocketDiags(t []netConnectionKindType, inodes map[string][]inodeMap) (map[string][]connTmp, error) {
//func getDockerSocketDiags(inodes map[string][]inodeMap) (map[string][]connTmp, error) {
//	waitgroup := &sync.WaitGroup{}
//	waitgroup.Add(1)
//
//	connSocketMap := make(map[string][]connTmp)
//	var err error = nil
//	var netNsModule *ns.NetNsMoudle = nil
//	go func() {
//		netNsModule, err = ns.NewNetNsModule()
//		if err != nil {
//			return
//		}
//		defer netNsModule.Close()
//
//		for _, inodeInfos := range inodes {
//			for _, inodeInfo := range inodeInfos {
//				_, ok := connSocketMap[inodeInfo.ns]
//				if ok == false {
//					err = netNsModule.SetNetNs(inodeInfo.pid)
//					if err == nil {
//						connSockets, err := processInetDiag()
//						if err == nil {
//							connSocketMap[inodeInfo.ns] = connSockets
//						}
//					}
//				}
//			}
//		}
//
//		waitgroup.Done()
//	}()
//
//	waitgroup.Wait()
//
//	return connSocketMap, err
//}

func ProcessInetDiag() ([]connTmp, error) {
	var ret []connTmp

	protos := []int{unix.IPPROTO_UDP, unix.IPPROTO_TCP}
	for _, proto := range protos {
		socketDiags, err := netlink.SocketDiag(syscall.AF_INET, uint8(proto))
		if err != nil {
			continue
		}

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
			} else if proto == unix.IPPROTO_UDP {
				sockType = syscall.SOCK_DGRAM
				if _, ok := UDPIntStatuses[state]; ok {
					if _, ok := UDPIntStatuses[state]; ok {
						status = UDPIntStatuses[state]
					}
				}

			}
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

		return ret, nil
	}
	return nil, nil
}
