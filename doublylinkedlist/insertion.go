package doublylinkedlist

import "fmt"

func InsertAtStart(head **DLNode, x int) {
	// We dont need to explicitly initialize Next & Prev. They will be given the zero value automatically.
	tmp := &DLNode{Data: x}
	fmt.Println("\nInserting", tmp, "at the start")
	if *head == nil {
		*head = tmp
		return
	}
	tmp.Next = *head
	(*head).Prev = tmp
	*head = tmp
}

func InsertAtEnd(head **DLNode, x int) {
	tmp := &DLNode{Data: x}
	fmt.Println("Inserting", tmp, "at the end")
	if *head == nil {
		*head = tmp
	} else {
		p := *head
		for ; p.Next != nil; p = p.Next {
		}
		p.Next = tmp
		tmp.Prev = p
	}
}

func InsertAtP(head **DLNode, p int, x int) {
	if p < 1 {
		fmt.Println("Positions start at 1. Invalid Position.")
		return
	}

	tmp := &DLNode{Data: x}
	if p == 1 {
		tmp.Next = *head
		if *head != nil {
			(*head).Prev = tmp
		}
		*head = tmp
		return
	}
	if *head == nil {
		fmt.Println("Position Invalid.")
		return
	}
	current := *head
	i := 1
	// Find the penultimate node
	for i < p-1 && current.Next != nil {
		current = current.Next
		i++
	}
	// This means we reached the end of the list and i != p-1
	if i < p-1 {
		fmt.Println("Invalid Position")
		return
	}
	fmt.Println("Inserting", tmp, "at position", p)
	tmp.Prev = current
	tmp.Next = current.Next
	if current.Next != nil {
		current.Next.Prev = tmp
	}
	current.Next = tmp
}
