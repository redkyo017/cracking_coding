package linkedlist

func DeleteMiddleNode(node *LinkedListNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return true
}
