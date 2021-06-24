package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func DisplayList(head *Node) {
	fmt.Printf("The list is: ")
	current := head
	for ; current != nil; current = current.Next {
		fmt.Print(current.Data, " ")
	}
	fmt.Println()
}

func Length(head *Node) int {
	current := head
	len := 0
	for ; current != nil; current = current.Next {
		len++
	}
	fmt.Println("Length = ", len)
	return len
}
