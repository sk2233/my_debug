package main

import (
	"fmt"
)

const CmdClear = "clear"

func RunClear(pid int, id int) {
	if id > 0 {
		clearByID(pid, id)
	} else {
		clearAll(pid)
	}
}

func clearAll(pid int) {
	for _, item := range breakPoints {
		WriteText(pid, item.Addr, item.Orig)
	}
	breakPoints = make(map[uint64]*BreakPoint)
}

func clearByID(pid int, id int) {
	var breakPoint *BreakPoint
	for _, item := range breakPoints {
		if item.ID == id {
			breakPoint = item
			break
		}
	}
	if breakPoint == nil {
		fmt.Printf("breakPoint %d not found\n", id)
		return
	}
	// 先恢复指令 再移除断点
	WriteText(pid, breakPoint.Addr, breakPoint.Orig)
	delete(breakPoints, breakPoint.Addr)
}
