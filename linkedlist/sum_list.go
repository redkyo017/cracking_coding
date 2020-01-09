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
