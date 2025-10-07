package main

import (
	"syscall"
)

const CmdAttach = "attach"

func RunAttach(pid int) {
	// 附加进程并等待进程的停止信号(并不是断点的停止信号)
	TraceAttach(pid)
	Wait4(pid, syscall.WSTOPPED)
	// 核心业务逻辑
	RunShell(pid)
	// 结束释放进程并清楚相关影响
	TraceDetach(pid)
	RunClear(pid, -1)
}
