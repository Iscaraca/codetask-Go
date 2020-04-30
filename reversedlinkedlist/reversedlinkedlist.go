package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	length int
	start  *node
}

func (list *linkedList) Append(newNode *node) {
	if list.length == 0 {
		list.start = newNode
	} else {
		currNode := list.start
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = newNode
	}
	list.length++
}

func (list *linkedList) Reverse() {
	curr := list.start
	following := list.start
	var prev *node = nil

	for curr != nil {
		following = following.next
		curr.next = prev
		prev = curr
		curr = following
	}

	list.start = prev
}

func main() {
	ll := &linkedList{}
	n1 := node{
		value: 1,
	}
	n2 := node{
		value: 5,
	}
	n3 := node{
		value: 9,
	}

	ll.Append(&n1)
	ll.Append(&n2)
	ll.Append(&n3)
	fmt.Printf("Initial order: %d, %d, %d\n", ll.start.value, ll.start.next.value, ll.start.next.next.value)
	// Prints 1, 5, 9

	ll.Reverse()

	fmt.Printf("Resultant order: %d, %d, %d\n", ll.start.value, ll.start.next.value, ll.start.next.next.value)
	// Prints 9, 5, 1
}
