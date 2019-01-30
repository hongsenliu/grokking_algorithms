package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(quickSort([]int{-317, -381, -14, -215, -590, -243, -412, 380, -312, 925, 158, -46, 177, 22, -482, 273, 217, 514, -392, 424}))
	ts := "2019-01-30T08:27:17.000-06:00"
	t, _ := time.Parse(time.RFC3339, ts)
	fmt.Println(t.Format(time.RFC3339))
}

func quickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	p := rand.Int() % len(items)
	left, right := 0, len(items)-1
	items[p], items[right] = items[right], items[p]

	for i := range items {
		if items[i] < items[right] {
			items[i], items[left] = items[left], items[i]
			left++
		}
	}
	items[left], items[right] = items[right], items[left]
	quickSort(items[:left])
	quickSort(items[left+1:])
	return items

}
