package linkedlist

import (
	"fmt"
)

func RunDemo() {
	var p *Node
	p = nil

	// Linked List Insertion
	InsertAtEndInPlace(&p, 2020)
	fmt.Printf("Adderss of node p %p\n", &*p)
	fmt.Println("Address of pointer p ", &p)
	DisplayList(p)

	p = InsertAtEnd(p, 1234)
	fmt.Printf("Adderss of node p %p\n", &*p)
	fmt.Println("Address of pointer p ", &p)
	DisplayList(p)

	InsertAtP(&p, 1, 666)
	DisplayList(p)

	InsertAtStart(&p, 4654654)
	DisplayList(p)

	InsertAtEndInPlace(&p, 1212)
	InsertAtEndInPlace(&p, 2323)
	InsertAtEndInPlace(&p, 3434)
	DisplayList(p)
	// Linked List Deletion
	p = DeleteFirst(p)
	DisplayList(p)

	DeleteFirstInPlace(&p)
	DisplayList(p)

	DeleteLast(&p)
	DisplayList(p)

	DeleteAtP(&p, 3)
	DisplayList(p)

	DeleteList(&p)
	DisplayList(p)

	InsertAtStart(&p, 1)
	DisplayList(p)
	DeleteAtP(&p, 2)
	DisplayList(p)
}
