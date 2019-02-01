package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(quickSort([]int{7, 2, 1, 8, 6, 3, 5, 4}))
	ts := "2019-01-30T08:27:17.000-06:00"
	t, _ := time.Parse(time.RFC3339, ts)
	fmt.Println(t.Format(time.RFC3339))
}

func quickSort(items []int) []int {
	if len(items) < 2 {
		fmt.Println("return:", items)
		return items
	}
	//p := rand.Int() % len(items)
	p := len(items) / 2
	left, right := 0, len(items)-1
	items[p], items[right] = items[right], items[p]

	for i := range items {
		if items[i] < items[right] {
			items[i], items[left] = items[left], items[i]
			left++
		}
	}
	items[left], items[right] = items[right], items[left]
	fmt.Println("left:", items[:left])
	quickSort(items[:left])
	//fmt.Println(len(items[left+1:]))
	fmt.Println("right:", items[left+1:])
	quickSort(items[left+1:])
	fmt.Println("RETURN:", items)
	return items

}
