package doublylinkedlist

import (
	"fmt"
)

func DeleteFirst(head **DLNode) {
	if *head == nil {
		fmt.Println("\nList empty")
		return
	}
	fmt.Println("\nDeleting first Node: ", *head)
	// We change head to point to the next Node
	*head = (**head).Next
	// To remove all references to the deleted node, set the Prev pointer of head to nil
	if *head != nil {
		(**head).Prev = nil
	}
}

func DeleteLast(head **DLNode) {
	if *head == nil {
		fmt.Println("\nList is empty")
		return
	}
	// If there is only 1 node in the list
	if (*head).Next == nil {
		*head = nil
		return
	}
	current := *head
	// Find the last node
	for current.Next != nil {
		current = current.Next
	}
	fmt.Println("\nDeleting last node: ", current)
	// Change the Next pointer of the penultimate node to nil
	current = current.Prev
	current.Next = nil
}

func DeleteAtP(head **DLNode, p int) {
	// Edge case 1: Positions less than 1
	if p < 1 {
		fmt.Println("Positions start from 1. Position Invaid.")
		return
	}
	// Edge case 2: List is empty
	if *head == nil {
		fmt.Println("List empty")
		return
	}
	// Edge case 3: Deletion position is 1
	if p == 1 {
		DeleteFirst(head)
		return
	}

	current := *head
	i := 1
	// Find node to delete
	for i < p && current.Next != nil {
		current = current.Next
		i++
	}
	// Edge case 4: Position exceeds length of the list
	if i < p {
		fmt.Println("Position Invalid")
		return
	}
	previousNode := current.Prev
	nextNode := current.Next
	previousNode.Next = nextNode
	// Edge case 5: The node to delete is the last in the list. If it is nextNode will be nil
	if nextNode != nil {
		nextNode.Prev = previousNode
	}
}

// DeleteList will delete the entire linkedlist.
// Deletion is inplace
func DeleteList(head **DLNode) {
	fmt.Println("\nDeleting entire linkedlist")
	var aux *DLNode
	for (*head) != nil {
		aux = (*head).Next
		*head = nil
		*head = aux
	}
}
