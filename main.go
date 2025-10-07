package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// https://www.hitzhangjie.pro/debugger101.io/

func main() {
	runtime.LockOSThread() // 需要保证后续操作的都是同一个线程
	cmd := os.Args[1]
	switch cmd {
	case CmdExec:
		RunExec(os.Args[2])
	case CmdAttach:
		pid, err := strconv.ParseInt(os.Args[2], 10, 64)
		HandleErr(err)
		RunAttach(int(pid))
	default:
		panic(fmt.Sprintf("unknown cmd %s", cmd))
	}
}
