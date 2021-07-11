package stack

import "fmt"

type DStack struct {
	capacity int
	top      int
	data     []int
}

func CreateDynStack() *DStack {
	newStack := &DStack{top: -1, data: make([]int, 0)}
	return newStack
}

func (s *DStack) isEmpty() bool {
	return s.top == -1
}

func (s *DStack) isFull() bool {
	return s.top == s.capacity-1
}

func (s *DStack) push(x int) {
	fmt.Println("\t\t STACK:", *s)
	fmt.Println("Pushing", x, "to the stack.")
	// If stack is full append the new element. This doubles the capacity of the stack. Then print the new capacity.
	if s.isFull() {
		(*s).data = append((*s).data, x)
		(*s).capacity = cap((*s).data)
		(*s).top++
		fmt.Println("\tStack Overflow. Increasing stack size to", (*s).capacity)
		return
	}
	(*s).data = append((*s).data, x)
	(*s).top++
}

func (s *DStack) pop() int {
	if s.isEmpty() {
		fmt.Println("Stack Underflow.")
		return -1
	}
	val := (*s).data[(*s).top]
	(*s).top--
	return val
}

func (s *DStack) peek() int {
	if s.isEmpty() {
		fmt.Print("Underflow.")
		return -1
	}
	return (*s).data[(*s).top]
}
