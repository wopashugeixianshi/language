package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
    finfo, _ := os.Stat("./pid.go")
    // Sys()返回的是interface{}，所以需要类型断言，不同平台需要的类型不一样，linux上为*syscall.Stat_t
    stat_t := finfo.Sys().(*syscall.Stat_t)
    fmt.Println(stat_t)
    // atime，ctime，mtime分别是访问时间，创建时间和修改时间，具体参见man 2 stat
    fmt.Println("atime:", timespecToTime(stat_t.Atim))
    fmt.Println("ctime:", timespecToTime(stat_t.Ctim))
    fmt.Println("mtime:", timespecToTime(stat_t.Mtim))
}

func timespecToTime(ts syscall.Timespec) time.Time {
    return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}
