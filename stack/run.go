package stack

import (
	"fmt"
)

func RunAStackDemo() {
	fmt.Println("---------- STATIC STACKS ----------")
	stack := CreateArrStack(5)
	for i := 10; i < 20; i++ {
		stack.push(i)
	}
	fmt.Println("Top of the stack =", stack.peek())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println(stack.peek())
}

func RunDStackDemo() {
	fmt.Println("---------- DYNAMIC STACKS ----------")
	stack := CreateDynStack()
	for i := 10; i < 20; i++ {
		stack.push(i)
	}
	fmt.Println("Top of the stack =", stack.peek())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println("Popped from the stack =", stack.pop())
	fmt.Println(stack.peek())
	for i := 1; i < 8; i++ {
		fmt.Println("Popped from the stack =", stack.pop())
	}
}
