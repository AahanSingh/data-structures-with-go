package linkedlist

import (
	"fmt"
)

/*
Go is garbage collected.
Variables that are unreachable are deleted by the garbage collector.
Thus, any pointers to the struct must not reference it anymore.
For illustration purposes, we will make the struct store its zero value as a replacement for the free() function from C.
Note that even after making the struct store its zero value, it will not actually deleted from memory.
The deletion will only be done once there are no pointers pointing to the struct thus making it unreachable.
Only after this will the garbage collector free memory used by the struct.
Another point to note is that once the garbage collector deletes the struct, it might now be available to the OS.
Go runtime will preiodically return large unused chunks of memory to the OS. This behavior can be forced through the use of the runtime.debug.FreeOSMemory() function.
*/

// DeleteFirst deletes the first node in a linked list.
// Deletion is not inplace. The function returns the new position of the head node after deletion.
func DeleteFirst(head *Node) *Node {
	if head == nil {
		fmt.Println("\nList empty")
		return head
	}
	fmt.Println("\nDeleting first Node: ", *head)
	// Create a new temporary pointer to point to the node being deleted.
	node := head
	// Head is pointing to the node. We must change head to make the node unreachable.
	head = head.Next
	// Make the struct store its zero value to illustrate that we are deleting the node.
	// Note that doing this will not actually delete the node.
	node.Next = nil
	node.Data = 0
	// Make the node unreachable so that the go garbage collector deletes this node.
	node = nil
	return head
}

// DeleteFirstInPlace deletes the first node in a linkedlist.
// Deletion is inplace.
func DeleteFirstInPlace(head **Node) {
	if *head == nil {
		fmt.Println("\nList empty")
		return
	}
	fmt.Println("\nDeleting first Node: ", *head)
	*head = (**head).Next
}

// DeleteAtP deletes the node located at index loction 'p'.
// Deletion is inplace.
func DeleteAtP(head **Node, p int) {
	if p == 0 {
		fmt.Println("\nPositions start from 1. Cannot delete.")
		return
	}
	if *head == nil {
		fmt.Println("\nList is empty")
	}
	// First node will be deleted
	if p == 1 {
		fmt.Println("\nDeleting Node: ", **head, "at position", p)
		*head = (*head).Next
		return
	}
	i := 1
	previous := &Node{}
	current := *head
	for current != nil && i < p {
		previous = current
		current = current.Next
		i++
	}
	if current == nil {
		fmt.Println("\nPosition does not exist")
		return
	}
	fmt.Println("\nDeleting Node: ", *current, "at position", p)
	previous.Next = current.Next
	current = nil
}

// DeleteLast deletes the last node in the linkedlist.
// Deletion is inplace
func DeleteLast(head **Node) {
	if *head == nil {
		fmt.Println("\nList is empty")
		return
	}
	current := (*head).Next
	previous := *head
	for ; current.Next != nil; current = current.Next {
		previous = current
	}
	previous.Next = nil
	fmt.Println("\nDeleting last node: ", current)
}

// DeleteList will delete the entire linkedlist.
// Deletion is inplace
func DeleteList(head **Node) {
	fmt.Println("\nDeleting entire linkedlist")
	var aux *Node
	for (*head) != nil {
		aux = (*head).Next
		*head = nil
		*head = aux
	}
}
