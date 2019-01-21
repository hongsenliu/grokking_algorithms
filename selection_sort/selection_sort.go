package main

import (
	"log"
)

func main() {
	d := []int{2, 1, 7, 5, 3, 6}
	selectionSort(d)
	log.Println(d)
}

func selectionSort(d []int) {
	l := len(d)
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i; j < l; j++ {
			if d[j] < d[minIndex] {
				minIndex = j
			}
		}
		d[i], d[minIndex] = d[minIndex], d[i]
	}
}
