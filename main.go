package main

import (
	// "cracking_coding/arraystrings"
	"cracking_coding/linkedlist"
	"log"
)

func main() {
	node1 := linkedlist.LinkedListNode{Val: 1, Next: nil}
	node2 := linkedlist.LinkedListNode{Val: 2, Next: &node1}
	node3 := linkedlist.LinkedListNode{Val: 3, Next: &node2}
	node4 := linkedlist.LinkedListNode{Val: 4, Next: &node3}
	node5 := linkedlist.LinkedListNode{Val: 5, Next: &node4}
	node6 := linkedlist.LinkedListNode{Val: 6, Next: &node5}
	// node7 := linkedlist.LinkedListNode{Val: 7, Next: nil}
	// node8 := linkedlist.LinkedListNode{Val: 8, Next: &node7}
	node9 := linkedlist.LinkedListNode{Val: 9, Next: &node3}
	node10 := linkedlist.LinkedListNode{Val: 10, Next: &node9}
	node11 := linkedlist.LinkedListNode{Val: 11, Next: &node10}
	node12 := linkedlist.LinkedListNode{Val: 12, Next: &node11}

	// log.Println(linkedlist.FindIntersection(&node12, &node6))
	n := linkedlist.FindIntersection(&node12, &node6)
	t := n
	for t != nil {
		log.Println("intersect node", t)
		t = t.Next
	}
}
