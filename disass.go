package main

import (
	"fmt"

	"golang.org/x/arch/x86/x86asm"
)

const CmdDisass = "disass"

func RunDisass(pid int, count int) {
	// 读取寄存器值
	regs := GetRegs(pid)
	// 若是因为 0xCC int3 产生的中断需要移动 pc 寄存器
	buff := ReadText(pid, regs.PC()-1, 1)
	if buff[0] == 0xCC { // TODO 这里因该恢复内容的
		regs.SetPC(regs.PC() - 1)
	}
	// 反汇编
	buff = ReadText(pid, regs.PC(), count)
	offset := 0
	for offset < len(buff) {
		// 使用64位模式
		inst, err := x86asm.Decode(buff[offset:], 64)
		HandleErr(err)
		fmt.Printf("%8x %s\n", uint64(offset)+regs.PC(), inst.String())
		offset += inst.Len
	}
}
