package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (linkedList *LinkedList) Insert(data int) {
	newNode := &Node{value: data, next: nil}

	if linkedList.head == nil {
		linkedList.head = newNode
	} else {
		current := linkedList.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (linkedList *LinkedList) Display() {
	current := linkedList.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	linkedList := LinkedList{}
	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Display()
}
