package doublylinkedlist

import "fmt"

type DLNode struct {
	Data int
	Prev *DLNode
	Next *DLNode
}

func DisplayList(head *DLNode) {
	fmt.Printf("The list is: ")
	for ; head != nil; head = head.Next {
		fmt.Print(head.Data, " ")
	}
	fmt.Println()
}

func Length(head *DLNode) int {
	len := 0
	for ; head != nil; head = head.Next {
		len++
	}
	fmt.Println("Length = ", len)
	return len
}
