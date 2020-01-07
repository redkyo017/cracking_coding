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

	log.Println(linkedlist.DeleteMiddleNode(&node5))
	// t := &node6
	// for t != nil {
	// 	log.Println("after delete", t)
	// 	t = t.Next
	// }
}
