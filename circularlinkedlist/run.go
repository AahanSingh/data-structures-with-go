package circularlinkedlist

func RunDemo() {
	var p *Node = nil

	InsertAtP(&p, 5, 4)
	DisplayList(p)

	InsertAtP(&p, 5, 0)
	DisplayList(p)

	InsertAtStart(&p, 1)
	DisplayList(p)

	InsertAtEnd(&p, 2)
	DisplayList(p)

	InsertAtEnd(&p, 3)
	DisplayList(p)

	InsertAtP(&p, 4, 4)
	DisplayList(p)

	DeleteFirst(&p)
	DisplayList(p)

	DeleteLast(&p)
	DisplayList(p)

	DeleteAtP(&p, 2)
	DisplayList(p)
	DeleteAtP(&p, 1)
	DisplayList(p)

	DeleteAtP(&p, 1)
	DisplayList(p)

	DeleteAtP(&p, 2)
	DisplayList(p)

	DeleteAtP(&p, 1)
	DisplayList(p)

	InsertAtEnd(&p, 1234)
	DisplayList(p)

	InsertAtP(&p, 3, 666)
	DisplayList(p)

	InsertAtStart(&p, 4654654)
	DisplayList(p)

	InsertAtEnd(&p, 1212)
	InsertAtEnd(&p, 2323)
	InsertAtEnd(&p, 3434)
	DisplayList(p)

	DeleteFirst(&p)
	DisplayList(p)

	DeleteLast(&p)
	DisplayList(p)

	DeleteAtP(&p, 3)
	DisplayList(p)

	DeleteList(&p)
	DisplayList(p)

}
