package main

import "fmt"

func main() {
	items := []int{}
	fmt.Println(lenSlice(items))
}

func lenSlice(items []int) int {
	if len(items) > 0 {
		return 1 + lenSlice(items[1:])
	}
	return 0
}
