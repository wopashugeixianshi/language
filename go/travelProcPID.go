package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	allPid, err := readPidsFromDir("/proc")
	if err != nil {
		log.Println("read all pid err:", err)
	}
	i := 0
	for k, v := range allPid {
		fmt.Println(i, k, v)
		i++
	}
}

func readPidsFromDir(path string) (map[int]string, error) {
	ret := make(map[int]string)
	var f *os.File
	d, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	fnames, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	for _, fname := range fnames {
		pid, err := strconv.ParseInt(fname, 10, 32)
		if err != nil {
			// if not numeric name, just skip
			continue
		}
		f, err = os.Open("/proc/" + strconv.Itoa(int(pid)) + "/status")
		if err != nil {
			log.Println("open err :", err)
		}
		defer f.Close()
		var ppidStr string
		readLine(f, func(line string) error {
			info := strings.SplitN(line, ":", 2)
			if len(info) == 2 {
				if info[0] == "PPid" {
					ppidStr = strings.TrimSpace(info[1])
					return fmt.Errorf("break readline, we just want ppid")
				}
			}
			return nil
		})
		ret[int(pid)] = ppidStr
	}

	return ret, nil
}

func readLine(reader io.Reader, f func(string) error) error {
	buf := bufio.NewReader(reader)
	line, err := buf.ReadBytes('\n')
	//l.Debug("line:", string(line))
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
