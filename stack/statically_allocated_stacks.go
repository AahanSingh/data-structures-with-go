package stack

import "fmt"

type AStack struct {
	capacity int
	top      int
	data     []int
}

func CreateArrStack(stack_size int) *AStack {
	newStack := &AStack{top: -1, capacity: stack_size, data: make([]int, stack_size)}
	return newStack
}

func (s *AStack) isEmpty() bool {
	return s.top == -1
}

func (s *AStack) isFull() bool {
	return s.top == s.capacity-1
}

func (s *AStack) push(x int) {
	fmt.Println("Pushing", x, "to the stack.")
	if s.isFull() {
		fmt.Println("Stack Overflow")
		return
	}
	(*s).top++
	(*s).data[(*s).top] = x
}

func (s *AStack) pop() int {
	if s.isEmpty() {
		fmt.Println("Stack Underflow.")
		return -1
	}
	val := (*s).data[(*s).top]
	(*s).top--
	return val
}

func (s *AStack) peek() int {
	if s.isEmpty() {
		fmt.Print("Underflow.")
		return -1
	}
	return (*s).data[(*s).top]
}
