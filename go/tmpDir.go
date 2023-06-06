package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	tmpDir, ioerr := ioutil.TempDir(os.TempDir(), "rpm")
	fmt.Println(tmpDir, ioerr)
	//os.RemoveAll(tmpDir)
	root := os.TempDir()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		//if err != nil {
		//	fmt.Println("who :", path, err)
		//	return err
		//}
		if err == nil && info.IsDir() && strings.Contains(path, "rpm") {
			fmt.Printf("Removing directory: %s\n", path)
			return os.RemoveAll(path)
		}
		return nil
	})
}
