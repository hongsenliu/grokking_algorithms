package main

import (
	"fmt"
)

var p *int

func main() {
	fmt.Println(max([]int{9}))
}
func max(items []int) int {
	if len(items) == 1 {
		return items[0]
	}
	if m := max(items[1:]); m >= items[0] {
		return m
	}
	return items[0]
}
