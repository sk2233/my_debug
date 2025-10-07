package main

import "syscall"

const CmdStep = "step"

func RunStep(pid int) {
	// 获取 pc-1 处的值
	regs := GetRegs(pid)
	buff := ReadText(pid, regs.PC()-1, 1)
	breakPoint := breakPoints[regs.PC()-1]
	// 对应位置的 int3 是打断点导致的需要进行恢复
	if buff[0] == 0xCC && breakPoint != nil {
		WriteText(pid, breakPoint.Addr, breakPoint.Orig)
		defer func() { // 执行完还要恢复
			WriteText(pid, breakPoint.Addr, []byte{0xCC})
		}()
		regs.SetPC(regs.PC() - 1) // 还要调整 pc 指针
		SetRegs(pid, regs)
	}
	// 进行单步调用并等待完成
	SingleStep(pid)
	Wait4(pid, syscall.WALL)
}
