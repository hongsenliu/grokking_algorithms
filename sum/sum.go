package main

import "fmt"

func main() {
	items := []int{-1, 1}
	fmt.Println(sum(items))
}

func sum(items []int) int {
	if len(items) > 0 {
		return items[len(items)-1] + sum(items[:len(items)-1])
	}
	return 0
}
