package main

import "syscall"

func GetRegs(pid int) *syscall.PtraceRegs {
	regs := &syscall.PtraceRegs{}
	err := syscall.PtraceGetRegs(pid, regs)
	HandleErr(err)
	return regs
}

func SetRegs(pid int, regs *syscall.PtraceRegs) {
	err := syscall.PtraceSetRegs(pid, regs)
	HandleErr(err)
}

func ReadText(pid int, pc uint64, size int) []byte {
	buff := make([]byte, size)
	count, err := syscall.PtracePeekText(pid, uintptr(pc), buff)
	HandleErr(err)
	return buff[:count]
}

func WriteText(pid int, addr uint64, buff []byte) {
	_, err := syscall.PtracePokeText(pid, uintptr(addr), buff)
	HandleErr(err)
}

func ReadData(pid int, pc uint64, size int) []byte {
	buff := make([]byte, size)
	count, err := syscall.PtracePeekData(pid, uintptr(pc), buff)
	HandleErr(err)
	return buff[:count]
}

func SingleStep(pid int) {
	err := syscall.PtraceSingleStep(pid)
	HandleErr(err)
}

func Wait4(pid int, options int) {
	var status syscall.WaitStatus
	var rusage syscall.Rusage
	_, err := syscall.Wait4(pid, &status, options, &rusage)
	HandleErr(err)
}

func TraceAttach(pid int) {
	err := syscall.PtraceAttach(pid)
	HandleErr(err)
}

func TraceDetach(pid int) {
	err := syscall.PtraceDetach(pid)
	HandleErr(err)
}

func Kill(pid int) {
	err := syscall.Kill(pid, syscall.SIGKILL)
	HandleErr(err)
}

func TraceCont(pid int) {
	err := syscall.PtraceCont(pid, 0)
	HandleErr(err)
}
