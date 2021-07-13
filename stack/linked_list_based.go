package stack

import (
	"fmt"

	"github.com/aahansingh/data-structures-with-go/linkedlist"
)

type stack struct {
	top       *linkedlist.Node
	stackSize int
}

func CreateStack() *stack {
	newStack := &stack{}
	return newStack
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func (s *stack) push(x int) {
	linkedlist.InsertAtStart(&s.top, x)
	s.stackSize++
}

func (s *stack) pop() int {
	if s.isEmpty() {
		fmt.Println("Stack Underflow.")
		return -1
	}
	val := (*s).top.Data
	linkedlist.DeleteFirstInPlace(&s.top)
	s.stackSize--
	return val
}

func (s *stack) peek() int {
	if s.isEmpty() {
		fmt.Print("Underflow.")
		return -1
	}
	return s.top.Data
}
