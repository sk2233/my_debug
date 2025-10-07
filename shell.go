package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const CmdExit = "exit"

func RunShell(pid int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		HandleErr(err)
		items := strings.Split(line, " ")
		args := make([]string, 0)
		for _, item := range items {
			item = strings.TrimSpace(item)
			if len(item) > 0 {
				args = append(args, item)
			}
		}
		switch args[0] {
		case CmdDisass:
			count := int64(64) // 默认 64 个可以指定
			if len(args) > 1 {
				count, err = strconv.ParseInt(args[1], 10, 64)
				HandleErr(err)
			}
			RunDisass(pid, int(count))
		case CmdBreak:
			addr, err := strconv.ParseUint(args[1], 16, 64)
			HandleErr(err)
			RunBreak(pid, addr)
		case CmdBreaks:
			RunBreaks()
		case CmdClear:
			id := int64(-1)
			if len(args) > 1 {
				id, err = strconv.ParseInt(args[1], 10, 64)
				HandleErr(err)
			}
			RunClear(pid, int(id))
		case CmdStep:
			RunStep(pid)
		case CmdContinue:
			RunContinue(pid)
		case CmdMem:
			addr, err := strconv.ParseUint(args[1], 16, 64)
			HandleErr(err)
			count, err := strconv.ParseInt(args[2], 10, 64)
			HandleErr(err)
			size, err := strconv.ParseInt(args[3], 10, 64)
			HandleErr(err)
			RunMem(pid, addr, int(count), int(size))
		case CmdRegs:
			RunRegs(pid)
		case CmdExit:
			return
		default:
			panic(fmt.Sprintf("unknown command: %s", args[0]))
		}
	}
}
