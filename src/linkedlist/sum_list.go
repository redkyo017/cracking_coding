package linkedlist

func SumList(l1 *LinkedListNode, l2 *LinkedListNode, carry int) *LinkedListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	var result LinkedListNode
	value := carry
	var l1Next, l2Next *LinkedListNode
	if l1 != nil {
		value += l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		value += l2.Val
		l2Next = l2.Next
	}
	result.Val = value % 10
	newCarry := 0
	if value >= 10 {
		newCarry = 1
	}
	if l1 != nil || l2 != nil {
		result.Next = SumList(l1Next, l2Next, newCarry)
	}
	return &result
}

func SumlistFollowUp(l1 *LinkedListNode, l2 *LinkedListNode) *LinkedListNode {
	len1 := Length(l1)
	len2 := Length(l2)
	if len1 < len2 {
		l1 = PadList(l1, len2-len1)
	} else {
		l2 = PadList(l2, len1-len2)
	}
	result, carry := SumListUpHelper(l1, l2)
	if carry != 0 {
		result = InsertBefore(result, carry)
	}
	return result
}

func SumListUpHelper(l1 *LinkedListNode, l2 *LinkedListNode) (result *LinkedListNode, carry int) {
	if l1 == nil && l2 == nil {
		return nil, 0
	}
	result, carry = SumListUpHelper(l1.Next, l2.Next)
	val := carry + l1.Val + l2.Val
	result = InsertBefore(result, val%10)
	return result, val / 10
}

func Length(node *LinkedListNode) (length int) {
	if node.Next == nil {
		return 1
	}
	p := node
	for p != nil {
		length++
		p = p.Next
	}
	return
}

func InsertBefore(node *LinkedListNode, val int) *LinkedListNode {
	var head LinkedListNode
	head.Val = val
	if node != nil {
		head.Next = node
	}
	return &head
}

func PadList(node *LinkedListNode, pad int) *LinkedListNode {
	head := node
	for i := 0; i < pad; i++ {
		head = InsertBefore(head, 0)
	}
	return head
}
