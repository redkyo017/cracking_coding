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
	// // solution 1: reversed and compare
	// reversed := ReverseAndClone(node)
	// return IsEqual(node, reversed)

	// // solution 2: stack , slow and fast pointer
	// slow := node
	// fast := node
	// stack := []int{}

	// for fast != nil && fast.Next != nil {
	// 	stack = append(stack, slow.Val)
	// 	slow = slow.Next
	// 	fast = fast.Next.Next
	// }

	// if fast != nil {
	// 	slow = slow.Next
	// }

	// for slow != nil {
	// 	top := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	if top != slow.Val {
	// 		return false
	// 	}
	// 	slow = slow.Next
	// }

	// solution 3: recursive
	listLength := Length(node)
	_, isPalindrome := IsPalindromeRecurse(node, listLength)
	return isPalindrome
	// return true
}

func IsPalindromeRecurse(node *LinkedListNode, length int) (*LinkedListNode, bool) {
	if node == nil || length <= 0 {
		return node, true
	} else if length == 1 {
		return node.Next, true
	}

	res, b := IsPalindromeRecurse(node.Next, length-2)
	if res == nil || b == false {
		return res, b
	}

	return res.Next, res.Val == node.Val
}
