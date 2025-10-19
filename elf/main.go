package main

import (
	"debug/dwarf"
	"debug/elf"
	"fmt"
	"os"
)

// https://www.hitzhangjie.pro/debugger101.io/

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
	//gosymtab, err := file.Section(".gosymtab").Data()
	//HandleErr(err)
	//gopclntab, err := file.Section(".gopclntab").Data()
	//HandleErr(err)
	//
	//pclntab := gosym.NewLineTable(gopclntab, file.Section(".text").Addr)
	//table, err := gosym.NewTable(gosymtab, pclntab)
	//HandleErr(err)
	//pc, _, err := table.LineToPC("/root/my_debug/test_main/main.go", 6)
	//HandleErr(err)
	//fmt.Println("PC:", pc)
	//fmt.Println(table.PCToLine(pc))
	data, err := file.DWARF()
	HandleErr(err)
	reader := data.Reader()
	for {
		entry, err := reader.Next()
		HandleErr(err)
		if entry == nil {
			break
		}
		// 编译对象，每个文件都是一个编译对象
		//if entry.Tag == dwarf.TagCompileUnit {
		//	temp, err := data.LineReader(entry)
		//	HandleErr(err)
		//	for _, item := range temp.Files() {
		//		if item == nil {
		//			continue
		//		}
		//		fmt.Println("CompileUnit", item.Name)
		//	}
		//}
		//// 编译方法
		//if entry.Tag == dwarf.TagSubprogram {
		//	fmt.Println("Subprogram", entry.Val(dwarf.AttrName))
		//}
		//// 变量
		//if entry.Tag == dwarf.TagVariable {
		//	fmt.Println("Variable", entry.Val(dwarf.AttrName))
		//}
		if entry.Tag == dwarf.TagCompileUnit {
			temp, err := data.LineReader(entry)
			HandleErr(err)
			item := dwarf.LineEntry{}
			err = temp.Next(&item)
			HandleErr(err)
			fmt.Println(item) // 获取行号信息
		}
	}
}
