package main

import "fmt"

func main() {
	fmt.Println(iteration([]int{2, 3, 4, 5, 7, 10}, 10))
	fmt.Println(recursive([]int{2, 3, 4, 5, 7, 10}, 190))
}

func iteration(s []int, target int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (low + high) / 2
		if s[mid] == target {
			return mid
		}
		if s[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func recursive(s []int, target int) int {
	return binarySearch(s, target, 0, len(s)-1)
}

func binarySearch(s []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if s[mid] == target {
		return mid
	}
	if s[mid] > target {
		return binarySearch(s, target, low, mid-1)
	}
	return binarySearch(s, target, mid+1, high)

}
