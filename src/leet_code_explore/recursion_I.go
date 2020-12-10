package leet_code_explore

import "log"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	reverseStringRecur(&s, i, j)

}
func reverseStringRecur(s *[]byte, i int, j int) {
	if j <= i {
		return
	}
	temp := (*s)[i]
	(*s)[i] = (*s)[j]
	(*s)[j] = temp
	i += 1
	j -= 1
	reverseStringRecur(s, i, j)
}

func reverseStringOpt(s []byte) {
	i := 0
	j := len(s) - 1
	for i <= j {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i += 1
		j -= 1
	}
}

func ReverseStringImplement() {
	s := "hello"
	chars := []byte(s)
	log.Println("con co ", chars)
	reverseStringOpt(chars)
	log.Println("con heo", chars)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}

func SwapPairsImplement() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	res := swapPairs(&node1)
	for res != nil {
		log.Println("con co", res)
		res = res.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	// var newRevertedNode *ListNode
	// for head != nil {
	// 	var newNode ListNode
	// 	newNode.Val = head.Val
	// 	newNode.Next = newRevertedNode
	// 	newRevertedNode = &newNode
	// 	head = head.Next
	// }
	// return newRevertedNode
	if head.Next == nil {
		return head
	}
	var newRevertedNode *ListNode
	node := reverseList(head.Next)
	node.Next = head

	newRevertedNode = node
	return newRevertedNode
}

func ReverseListEmplement() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node5 := ListNode{5, nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	node := reverseList(&node1)
	for node != nil {
		log.Println("con co", node)
		node = node.Next
	}
}
