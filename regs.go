package main

import (
	"fmt"
	"reflect"
	"sort"
)

const CmdRegs = "regs"

type RegItem struct {
	Name  string
	Value uint64
}

func RunRegs(pid int) {
	regs := GetRegs(pid) // 这里展示的与 linux 下的源码不一样
	rt := reflect.TypeOf(regs).Elem()
	rv := reflect.ValueOf(regs).Elem()
	items := make([]*RegItem, 0)
	for i := 0; i < rv.NumField(); i++ { // 写入值是正好相反的操作
		items = append(items, &RegItem{
			Name:  rt.Field(i).Name,
			Value: rv.Field(i).Uint(),
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if len(items[i].Name) == len(items[j].Name) {
			return items[i].Name < items[j].Name
		}
		return len(items[i].Name) < len(items[j].Name)
	})
	for _, item := range items {
		fmt.Printf("Register\t%s\t%016x\n", item.Name, item.Value)
	}
}
