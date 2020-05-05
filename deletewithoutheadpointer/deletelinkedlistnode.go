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

func (list *linkedList) Print() {
	node := list.start
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Printf("\n")
}

// All it does is set the memory address of the deleted node to the node that comes next, so it skips the node essentially
func delete(node *node) {
	*node = *node.next
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
	n4 := node{
		value: 11,
	}
	n5 := node{
		value: 13,
	}

	ll.Append(&n1)
	ll.Append(&n2)
	ll.Append(&n3)
	ll.Append(&n4)
	ll.Append(&n5)

	ll.Print() // Prints 1 5 9 11 13 

	delete(&n3)

	ll.Print() // Prints 1 5 11 13
}
