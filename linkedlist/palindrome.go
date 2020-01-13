package linkedlist

func ReverseAndClone(node *LinkedListNode) *LinkedListNode {
	var newReveredNode *LinkedListNode
	for node != nil {
		var newNode LinkedListNode
		newNode.Val = node.Val
		newNode.Next = newReveredNode
		newReveredNode = &newNode
		node = node.Next
	}
	return newReveredNode
}

func IsEqual(l1 *LinkedListNode, l2 *LinkedListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func IsPalindrome(node *LinkedListNode) bool {
	// solution 1: reversed and compare
	reversed := ReverseAndClone(node)
	return IsEqual(node, reversed)
}
