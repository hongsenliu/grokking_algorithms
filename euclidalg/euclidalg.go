package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcd(1680, 640))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
