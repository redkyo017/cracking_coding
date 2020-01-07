package linkedlist

import "log"

func ReturnKthToLast(head LinkedListNode, k int) LinkedListNode {
	// r := &head
	// for r != nil {
	// 	log.Println("con co", r)
	// 	r = r.Next
	// }
	// log.Println(printKthToLast(&head, k))
	KthNode := returnKthToLastPointer(&head, k)
	log.Println("con co", KthNode)
	return *KthNode
}

func printKthToLast(head *LinkedListNode, k int) int {
	if head == nil {
		return 0
	}
	index := printKthToLast(head.Next, k) + 1
	if index == k {
		log.Println(k, "th to last node is ", head.Val)
	}
	return index
}

func returnKthToLast(head *LinkedListNode, k, i int) *LinkedListNode {
	if head == nil {
		return nil
	}
	if i == k {
		return head
	}
	i++
	node := returnKthToLast(head.Next, k, i)
	return node
}

func returnKthToLastPointer(head *LinkedListNode, k int) *LinkedListNode {
	p1 := head
	p2 := head

	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
