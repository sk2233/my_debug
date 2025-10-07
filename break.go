package main

const CmdBreak = "break"

type BreakPoint struct {
	ID   int
	Addr uint64
	Orig []byte
}

var (
	breakPoints  = make(map[uint64]*BreakPoint)
	breakPointID = 0
)

func GenBreakPointID() int {
	breakPointID++
	return breakPointID
}

func RunBreak(pid int, addr uint64) {
	// 先保存原内容
	buff := ReadText(pid, addr, 1)
	breakPoints[addr] = &BreakPoint{
		ID:   GenBreakPointID(),
		Addr: addr,
		Orig: buff,
	}
	// 再替换为中断指令
	WriteText(pid, addr, []byte{0xCC})
}
