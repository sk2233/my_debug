package main

import (
	"os/exec"
	"syscall"
)

const CmdExec = "exec"

func RunExec(prog string) {
	// 设置新进程启动后等待调试器连接
	cmd := exec.Command(prog)
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	err := cmd.Start()
	HandleErr(err)
	// 等待进程的所有事件
	Wait4(cmd.Process.Pid, syscall.WALL)
	// 核心业务逻辑
	RunShell(cmd.Process.Pid)
	// 结束释放并杀死进程
	TraceDetach(cmd.Process.Pid)
	Kill(cmd.Process.Pid)
}
