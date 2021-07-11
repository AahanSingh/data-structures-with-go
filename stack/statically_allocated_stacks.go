package stack

import "fmt"

type SStack struct {
	capacity int
	top      int
	data     []int
}

func CreateArrStack(stack_size int) *SStack {
	newStack := &SStack{top: -1, capacity: stack_size, data: make([]int, stack_size)}
	return newStack
}

func (s *SStack) isEmpty() bool {
	return s.top == -1
}

func (s *SStack) isFull() bool {
	return s.top == s.capacity-1
}

func (s *SStack) push(x int) {
	fmt.Println("Pushing", x, "to the stack.")
	if s.isFull() {
		fmt.Println("Stack Overflow")
		return
	}
	(*s).top++
	(*s).data[(*s).top] = x
}

func (s *SStack) pop() int {
	if s.isEmpty() {
		fmt.Println("Stack Underflow.")
		return -1
	}
	val := (*s).data[(*s).top]
	(*s).top--
	return val
}

func (s *SStack) peek() int {
	if s.isEmpty() {
		fmt.Print("Underflow.")
		return -1
	}
	return (*s).data[(*s).top]
}
