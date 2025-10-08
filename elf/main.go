package main

import (
	"debug/elf"
	"debug/gosym"
	"fmt"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := elf.Open(os.Args[1]) // 读取 elf 文件
	HandleErr(err)
	//for _, section := range open.Sections {
	//	fmt.Println(section.Name)
	//}
	//data, err := file.Section(".text").Data() // 读取代码段
	//HandleErr(err)
	//offset := 0
	//for i := 0; i < 10; i++ {
	//	inst, err := x86asm.Decode(data[offset:], 64)
	//	HandleErr(err)
	//	fmt.Println(inst.String())
	//	offset += inst.Len
	//}
	gosymtab, err := file.Section(".gosymtab").Data()
	HandleErr(err)
	gopclntab, err := file.Section(".gopclntab").Data()
	HandleErr(err)

	pclntab := gosym.NewLineTable(gopclntab, file.Section(".text").Addr)
	table, err := gosym.NewTable(gosymtab, pclntab)
	HandleErr(err)
	pc, _, err := table.LineToPC("/root/my_debug/test_main/main.go", 6)
	HandleErr(err)
	fmt.Println("PC:", pc)
	fmt.Println(table.PCToLine(pc))
}
