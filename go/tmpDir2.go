package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	"path/filepath"
	//"strings"
)

func main() {
	root := filepath.Join(os.TempDir(), "RPM")
	fmt.Println(root)
	err := os.RemoveAll(root)
	fmt.Println(err)
	err = os.MkdirAll(root, 0700)
	fmt.Println(err)
	_, err = os.Stat(root) 
	if err != nil {
		os.MkdirAll(root, 0700)
	}
	//tmpDir, err := ioutil.TempDir(root, "rpm")
	//fmt.Println(tmpDir, err)
	//_, _ = ioutil.TempDir(root, "rpm")
	//_, _ = ioutil.TempDir(root, "rpm")
	//_, _ = ioutil.TempDir(root, "rpm")
	//_, _ = ioutil.TempDir(root, "rpm")
	//_, _ = ioutil.TempDir(root, "rpm")
	//_, _ = ioutil.TempDir(root, "rpm")
	//_, _ = ioutil.TempDir(root, "rpm")
	//defer os.RemoveAll(root)
}
