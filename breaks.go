package main

import (
	"fmt"
	"sort"
)

const CmdBreaks = "breaks"

func RunBreaks() {
	items := make([]*BreakPoint, 0)
	for _, item := range breakPoints {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	for _, item := range items {
		fmt.Printf("breakpoint[%d] %#x\n", item.ID, item.Addr)
	}
}
