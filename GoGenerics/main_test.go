package main

import "testing"

func TestStack_pop(t *testing.T) {
	stack := Stack[int]{}

	_, found := stack.Pop()

	if found {
		t.Errorf("Error as stack data is empty")
	}

}
