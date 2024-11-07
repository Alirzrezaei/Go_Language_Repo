package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}
func (s *Stack[T]) Peak() (value T, found bool) {
	if len(s.data) == 0 {
		return value, false
	}
	return s.data[len(s.data)-1], true
}
func (s *Stack[T]) Pop() (T, bool) {
	value, found := s.Peak()
	if !found {
		return value, false
	}
	s.data = s.data[:len(s.data)-1]
	return value, true
}
func main() {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(25)

	fmt.Println("my stack is: ", stack.data)

	fmt.Println(stack.Peak())

	fmt.Println(stack.Pop())

	fmt.Println(stack.Peak())

}
