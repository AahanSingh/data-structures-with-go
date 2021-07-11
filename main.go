package main

import (
	"fmt"

	"github.com/aahansingh/data-structures-with-go/circularlinkedlist"
	"github.com/aahansingh/data-structures-with-go/doublylinkedlist"
	"github.com/aahansingh/data-structures-with-go/linkedlist"
	"github.com/aahansingh/data-structures-with-go/stack"
)

func main() {
	fmt.Println("---------- LINKED LIST ---------- ")
	linkedlist.RunDemo()
	fmt.Println("---------- DOUBLY LINKED LIST ---------- ")
	doublylinkedlist.RunDemo()
	fmt.Println("---------- CIRCULAR LINKED LIST ---------- ")
	circularlinkedlist.RunDemo()
	fmt.Println("---------- ARRAY BASED STACK ---------- ")
	stack.RunAStackDemo()
	stack.RunDStackDemo()
}
