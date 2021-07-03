package circularlinkedlist

import (
	"fmt"
)

func DeleteFirst(head **Node) {
	if *head == nil {
		fmt.Println("\nList empty")
		return
	}
	fmt.Println("\nDeleting first Node: ", *head)
	// Edge Case 1: Only one node exists in the list
	if (*head).Next == *head {
		*head = nil
		return
	}
	// Traverse to last node
	p := *head
	for ; p.Next != *head; p = p.Next {
	}
	p.Next = (*head).Next
	*head = (*head).Next
}

func DeleteLast(head **Node) {
	if *head == nil {
		fmt.Println("\nList is empty")
		return
	}
	// Edge Case 1: If there is only 1 node in the list
	if (*head).Next == *head {
		*head = nil
		return
	}
	current := (*head).Next
	previous := *head
	for current.Next != *head {
		previous = current
		current = current.Next
	}
	fmt.Println("\nDeleting last node: ", current)
	previous.Next = *head
	current.Next = nil
}

func DeleteAtP(head **Node, p int) {
	if p < 1 {
		fmt.Println("\nPositions start from 1. Cannot delete.")
		return
	}
	if *head == nil {
		fmt.Println("\nList is empty")
		return
	}
	if p == 1 {
		DeleteFirst(head)
		return
	}
	i := 1
	var previous *Node = nil
	current := *head
	for i < p && current.Next != *head {
		previous = current
		current = current.Next
		i++
	}
	if i < p {
		fmt.Println("Invalid position.")
		return
	}
	fmt.Println("Deleting", current, "at position", p)
	previous.Next = current.Next
	current = nil
}

func DeleteList(head **Node) {
	fmt.Println("\nDeleting entire linkedlist")
	var aux *Node
	for *head != nil {
		aux = (*head).Next
		if aux == *head || aux == nil {
			*head = nil
			return
		}
		(*head).Next = nil
		*head = aux
	}
}
