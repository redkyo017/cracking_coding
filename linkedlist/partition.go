package linkedlist

func Partition(node *LinkedListNode, value int) *LinkedListNode {
	// // solution 1: use 4 pointer
	// var beforeStart, beforeEnd, afterStart, afterEnd *LinkedListNode
	// for node != nil {
	// 	next := node.Next
	// 	node.Next = nil
	// 	if node.Val < value {
	// 		if beforeStart == nil {
	// 			beforeStart = node
	// 			beforeEnd = beforeStart
	// 		} else {
	// 			beforeEnd.Next = node
	// 			beforeEnd = node
	// 		}
	// 	} else {
	// 		if afterStart == nil {
	// 			afterStart = node
	// 			afterEnd = afterStart
	// 		} else {
	// 			afterEnd.Next = node
	// 			afterEnd = node
	// 		}
	// 	}
	// 	node = next
	// }
	// if beforeStart == nil {
	// 	return afterStart
	// }
	// beforeEnd.Next = afterStart
	// return beforeStart

	// solution 2: use 2 pointer
	head, tail := node, node
	for node != nil {
		next := node.Next
		if node.Val < value {
			node.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}
		node = next
	}
	tail.Next = nil
	// t := head
	// for t != nil {
	// 	log.Println("con co", t)
	// 	t = t.Next
	// }
	return head
}
