package main

import (
	"fmt"

	"github.com/aahansingh/data-structures-with-go/doublylinkedlist"
	"github.com/aahansingh/data-structures-with-go/linkedlist"
)

func main() {
	fmt.Println("---------- LINKED LIST ---------- ")
	linkedlist.RunDemo()
	fmt.Println("---------- DOUBLY LINKED LIST ---------- ")
	doublylinkedlist.RunDemo()
}
