package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	ppid := GetProcPPID("/proc/", 1931)
	fmt.Println("ppid :", ppid)
}

func GetProcPPID(procRoot string, pid uint) int {
	var ppid int
	f, err := os.Open(path.Join(procRoot, fmt.Sprintf("%d/status", pid)))
	if err != nil {
		return 0
	}
	defer f.Close()
	var count int
	ReadLine(f, func(line string) error {
		count++
		info := strings.SplitN(line, ":", 2)
		if len(info) == 2 {
			if info[0] == "PPid" {
				ppid, _ = strconv.Atoi(strings.TrimSpace(info[1]))
				return fmt.Errorf("break readline, we just want Name")
			}
		}
		return nil
		//return fmt.Errorf("break readline, we just want Name")
	})
	fmt.Println(count)
	return ppid
}

func ReadLine(reader io.Reader, f func(string) error) error {
	buf := bufio.NewReader(reader)
	line, err := buf.ReadBytes('\n')
	for err == nil {
		line = bytes.TrimRight(line, "\n")
		if len(line) > 0 {
			if line[len(line)-1] == 13 { //'\r'
				line = bytes.TrimRight(line, "\r")
			}
			err = f(string(line))
			if err != nil {
				return err
			}
		}
		line, err = buf.ReadBytes('\n')
	}
	if len(line) > 0 {
		err = f(string(line))
	}
	if err != io.EOF {
		return err
	}
	return nil
}
