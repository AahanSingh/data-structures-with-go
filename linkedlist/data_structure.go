package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func DisplayList(head *Node) {
	fmt.Printf("The list is: ")
	for ; head != nil; head = head.Next {
		fmt.Print(head.Data, " ")
	}
	fmt.Println()
}

func Length(head *Node) int {
	len := 0
	for ; head != nil; head = head.Next {
		len++
	}
	fmt.Println("Length = ", len)
	return len
}
