package linkedlist

import (
	"math"
)

func FindIntersection(l1 *LinkedListNode, l2 *LinkedListNode) *LinkedListNode {
	tail1, len1 := GetLengthAndTail(l1)
	tail2, len2 := GetLengthAndTail(l2)

	if tail1 != tail2 {
		return nil
	}
	var longer, shorter *LinkedListNode
	if len1 < len2 {
		longer = l2
		shorter = l1
	} else {
		longer = l1
		shorter = l2
	}
	nodeKth := GetNodeFromKth(longer, int(math.Abs(float64(len1)-float64(len2))))

	for nodeKth != shorter {
		nodeKth = nodeKth.Next
		shorter = shorter.Next
	}
	return nodeKth
}

func GetLengthAndTail(node *LinkedListNode) (*LinkedListNode, int) {
	if node == nil {
		return nil, 0
	}
	length := 1
	tail := node
	for tail.Next != nil {
		length++
		tail = tail.Next
	}
	return tail, length
}

func GetNodeFromKth(head *LinkedListNode, pos int) *LinkedListNode {
	if head == nil {
		return nil
	}
	node := head
	for pos != 0 {
		pos--
		node = node.Next
	}
	return node
}
