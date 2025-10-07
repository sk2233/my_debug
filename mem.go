package main

import (
	"encoding/binary"
	"fmt"
)

const CmdMem = "mem" // 写入是与读取正好相反

func RunMem(pid int, addr uint64, count int, size int) {
	buff := ReadData(pid, addr, count*size)
	if len(buff) != count*size {
		fmt.Printf("remain byte %d not enough\n", len(buff))
		return
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%s\t", formatNum(buff[i*size:], size))
	}
	fmt.Print("\n")
}

func formatNum(buff []byte, size int) string {
	switch size {
	case 2:
		return fmt.Sprintf("0x%04X", binary.LittleEndian.Uint16(buff))
	case 4:
		return fmt.Sprintf("0x%08X", binary.LittleEndian.Uint32(buff))
	case 8:
		return fmt.Sprintf("0x%016X", binary.LittleEndian.Uint64(buff))
	default: // 其他的都视为1个
		return fmt.Sprintf("0x%02X", buff[0])
	}
}
