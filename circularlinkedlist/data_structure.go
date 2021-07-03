package circularlinkedlist

import "fmt"

// This is the same as the struct for a singly linked list.
type Node struct {
	Data int
	Next *Node
}

func DisplayList(head *Node) {
	fmt.Printf("The list is: ")
	current := head
	if current == nil {
		fmt.Println("empty.")
		return
	} else {
		fmt.Print(" -> ", current)
		current = current.Next
		for ; current != head; current = current.Next {
			fmt.Print(" -> ", current)
		}
		fmt.Println()
		fmt.Println()
	}
}

func Length(head *Node) int {
	current := head
	if current == nil {
		return 0
	} else {
		len := 1
		current = current.Next
		for ; current != head; current = current.Next {
			len++
		}
		return len
	}
}
