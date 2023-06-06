package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	// "unsafe"
)

var (
	byteOrder = binary.LittleEndian
)

const (
	PROC_CN_MCAST_LISTEN = 1
	PROC_EVENT_NONE      = 0
	PROC_EVENT_FORK      = 1
	PROC_EVENT_EXEC      = 2
	PROC_EVENT_UID       = 3
	PROC_EVENT_GID       = 4
	PROC_EVENT_SID       = 5
)

type cbId struct {
	Idx uint32
	Val uint32
}

//type cnMsg struct {
//	hdr  syscall.NlMsghdr
//	data [1024]byte
//}

type procEventHeader struct {
	What      uint32
	Cpu       uint32
	Timestamp uint64
}

type cnMsg struct {
	Id    cbId
	Seq   uint32
	Ack   uint32
	Len   uint16
	Flags uint16
}

type forkProcEvent struct {
	ParentPid  uint32
	ParentTgid uint32
	ChildPid   uint32
	ChildTgid  uint32
}
type execProcEvent struct {
	ProcessPid  uint32
	ProcessTgid uint32
}

type exitProcEvent struct {
	ProcessPid  uint32
	ProcessTgid uint32
	ExitCode    uint32
	ExitSignal  uint32
}

func main() {
	// 创建套接字
	fd, _ := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, syscall.NETLINK_CONNECTOR)
	addr := &syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Groups: PROC_CN_MCAST_LISTEN,
	}
	// 绑定套接字
	syscall.Bind(fd, addr)

	// 读取进程事件消息
	buf := make([]byte, syscall.Getpagesize())
	for {
		n, _, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		msgs, _ := syscall.ParseNetlinkMessage(buf[:n])
		for _, m := range msgs {
			if m.Header.Type == syscall.NLMSG_DONE {
				handleEvent(m.Data)
			}
		}
		/*
			// 解析消息
			for i := uint32(0); i < uint32(n); {
				msg := (*cnMsg)(unsafe.Pointer(&buf[i]))
				i += msg.hdr.Len

				//if msg.hdr.Type == syscall.NLMSG_DONE {
				//	break
				//}

				if msg.hdr.Type == syscall.NLMSG_ERROR {
					fmt.Println("netlink error")
					return
				}

				msgs, _ := syscall.ParseNetlinkMessage(buf[:nr])
				// 获取进程 ID
				pid := *(*uint32)(unsafe.Pointer(&msg.data[0]))

				// 处理事件类型
				switch *(*uint32)(unsafe.Pointer(&msg.data[4])) {
				case PROC_EVENT_FORK:
					// 进程 fork 事件
					childPid := *(*uint32)(unsafe.Pointer(&msg.data[8]))
					fmt.Printf("process %d forked child process %d\n", pid, childPid)
				case PROC_EVENT_EXEC:
					// 进程 exec 事件
					fmt.Printf("process %d executed a new program\n", pid)
				case PROC_EVENT_UID:
					// 进程 uid 改变事件
					uid := *(*uint32)(unsafe.Pointer(&msg.data[8]))
					fmt.Printf("process %d uid changed to %d\n", pid, uid)
				case PROC_EVENT_GID:
					// 进程 gid 改变事件
					gid := *(*uint32)(unsafe.Pointer(&msg.data[8]))
					fmt.Printf("process %d gid changed to %d\n", pid, gid)
				case PROC_EVENT_SID:
					// 进程 sid 改变事件
					sid := *(*uint32)(unsafe.Pointer(&msg.data[8]))
					fmt.Printf("process %d sid changed to %d\n", pid, sid)
				}
			}
		*/
	}
}

func handleEvent(data []byte) {
	buf := bytes.NewBuffer(data)
	msg := &cnMsg{}
	hdr := &procEventHeader{}

	binary.Read(buf, byteOrder, msg)
	binary.Read(buf, byteOrder, hdr)
	switch hdr.What {
	case PROC_EVENT_FORK:
		event := &forkProcEvent{}
		binary.Read(buf, byteOrder, event)
		//ppid := int(event.ParentTgid)
		//pid := int(event.ChildTgid)
		//fmt.Println("fork pid =", pid, " tgid =", ppid)
		fmt.Println(event)
	case PROC_EVENT_EXEC:
		event := &execProcEvent{}
		binary.Read(buf, byteOrder, event)
		pid := int(event.ProcessTgid)
		ppid := int(event.ProcessTgid)
		fmt.Println("exec pid =", pid, " tgid =", ppid)
		//case PROC_EVENT_EXIT:
		//	event := &exitProcEvent{}
		//	binary.Read(buf, byteOrder, event)
		//	pid := int(event.ProcessPid)
		//	ppid := int(event.ProcessTgid)
		//	fmt.Println("exit pid =", pid, " tgid =", ppid)
	}
}
