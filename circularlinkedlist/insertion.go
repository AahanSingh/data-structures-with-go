package circularlinkedlist

import "fmt"

func InsertAtStart(head **Node, x int) {
	// Create new node
	tmp := &Node{Data: x}
	tmp.Next = tmp

	fmt.Println("\nInserting", tmp, "at the start ")
	if *head == nil {
		*head = tmp
		return
	}
	p := *head
	for ; p.Next != *head; p = p.Next {
	}
	p.Next = tmp
	tmp.Next = *head
	*head = tmp
}

func InsertAtEnd(head **Node, x int) {
	// Create a new node & make it point to itself.
	tmp := &Node{Data: x}
	tmp.Next = tmp

	fmt.Println("\nInserting", tmp, "at the end")
	if *head == nil {
		*head = tmp
	} else {
		p := *head
		// Find the last node
		for ; p.Next != *head; p = p.Next {
		}
		p.Next = tmp
		// Make the new node point to head
		tmp.Next = *head
	}
}

func InsertAtP(head **Node, p int, x int) {
	if p < 1 {
		fmt.Println("\nPositions start at 1. Invalid input position.")
		return
	}

	if p == 1 {
		InsertAtStart(head, x)
		return
	}
	// If insertion point is not the start and the list is empty, insertion point is invalid.
	if *head == nil {
		fmt.Println("\nInvalid Position.")
		return
	}

	current := *head
	i := 1
	for i < p-1 && current.Next != *head {
		current = current.Next
		i++
	}
	if i < p-1 {
		fmt.Println("\nInvalid Position")
		return
	}
	tmp := &Node{Data: x}
	tmp.Next = tmp
	fmt.Println("\nInserting", tmp, "at the position", p)
	tmp.Next = current.Next
	current.Next = tmp
}
