package main

import (
	"errors"
	"fmt"
)

// Stack with slice
type Stack struct {
	items []int
}

func (s *Stack) push(v int) {
	if len(s.items) == 0 {
		s.items = append(s.items, v)
	}
}

func (s *Stack) pop() (int, error) {
	if l := len(s.items); l > 0 {
		v := s.items[l-1]
		s.items = s.items[:l-1]
		return v, nil
	}
	return 0, errors.New("stack is empty")
}
func main() {
	s1 := Stack{}
	s1.push(1)
	s1.push(2)
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())
}
