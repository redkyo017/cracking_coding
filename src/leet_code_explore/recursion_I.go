package leet_code_explore

import (
	"log"
	"math"
)

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
	// SOLUTION 1: ITERATIVE
	// var newRevertedNode *ListNode
	// for head != nil {
	// 	var newNode ListNode
	// 	newNode.Val = head.Val
	// 	newNode.Next = newRevertedNode
	// 	newRevertedNode = &newNode
	// 	head = head.Next
	// }
	// return newRevertedNode
	// SOLUTION 2: RECURSIVE
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
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

// SEARCH IN A BINARY SEARCH TREE
// Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	nodeLeft := searchBST(root.Left, val)
	nodeRight := searchBST(root.Right, val)
	if nodeLeft == nil {
		return nodeRight
	}
	return nodeLeft
}

func SearchBSTimplement() {
	node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node4 := TreeNode{4, nil, nil}
	node5 := TreeNode{5, nil, nil}

	root := &node4
	root.Left = &node2
	root.Right = &node5
	node2.Left = &node1
	node2.Right = &node3

	searched := searchBST(root, 2)
	if searched != nil {
		log.Println(searched)
		log.Println(searched.Left)
		log.Println(searched.Right)
	}
}

// PASCAL'S TRIANGLE II
// Given an integer rowIndex, return the rowIndex-th row of the Pascal's triangle.
// Notice that the row index starts from 0.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.
// Follow up:
// Could you optimize your algorithm to use only O(k) extra space?
// Example 1:
// Input: rowIndex = 3
// Output: [1,3,3,1]

// Example 2:
// Input: rowIndex = 0
// Output: [1]

// Example 3:
// Input: rowIndex = 1
// Output: [1,1]
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	res := []int{}
	res = append(res, 1)
	prevRow := getRow(rowIndex - 1)
	for j := 0; j < len(prevRow)-1; j++ {
		res = append(res, prevRow[j]+prevRow[j+1])
	}
	res = append(res, 1)
	return res
}

func GetRowPascalTriangleImplement() {
	rowIndex := 5
	log.Println(getRow(rowIndex))
}

func fib(N int) int {
	if N < 2 {
		return N
	}
	return fib(N-1) + fib(N-2)
}
func fibMemozation(N int, memo map[int]int) int {
	if N < 2 {
		return N
	}
	if v, ok := memo[N]; ok {
		return v
	}
	memo[N] = fibMemozation(N-1, memo) + fibMemozation(N-2, memo)
	return memo[N]
}
func FibImplement() {
	memo := map[int]int{}
	log.Println("con co", fibMemozation(6, memo))
}

// CLIMBING STAIRS
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps

// Example 2:
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsMemo(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
	return memo[n]
}

func ClimbStairsImplement() {
	memo := map[int]int{}
	log.Println(climbStairsMemo(3, memo))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left_height := maxDepth(root.Left)
	right_height := maxDepth(root.Right)

	return int(math.Max(float64(left_height), float64(right_height))) + 1
}

func MaxDepthBinaryTreeSolution() {
	node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node4 := TreeNode{3, nil, nil}
	node5 := TreeNode{3, nil, nil}
	node1.Right = &node2
	node2.Left = &node3
	node3.Right = &node4
	node4.Left = &node5
	log.Println(maxDepth(&node1))
}

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if x == 0 {
		return 0
	}
	isNegativePow := false
	if n < 0 {
		isNegativePow = true
		n = int(math.Abs(float64(n)))
	}
	result := x * myPow(x, n-1)
	if isNegativePow {
		result = 1 / result
	}
	return result
}

func MyPowSolution() {
	log.Println("con co", myPow(2, -10))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node *ListNode
	if l1.Val < l2.Val {
		node = l1
		node.Next = mergeTwoLists(l1.Next, l2)
	} else {
		node = l2
		node.Next = mergeTwoLists(l1, l2.Next)
	}
	return node
}

func MergeTwoListsSolution() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node5 := ListNode{5, nil}
	node6 := ListNode{6, nil}
	node1.Next = &node3
	node3.Next = &node5
	node2.Next = &node4
	node4.Next = &node6

	mergeListNode := mergeTwoLists(&node1, &node2)
	node := mergeListNode
	for node != nil {
		log.Println("con co", node)
		node = node.Next
	}
}

// K-TH SYMBOL IN GRAMMAR
// On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.
// Given row N and index K, return the K-th indexed symbol in row N. (The values of K are 1-indexed.) (1 indexed).
// Examples:
// Input: N = 1, K = 1
// Output: 0

// Input: N = 2, K = 1
// Output: 0

// Input: N = 2, K = 2
// Output: 1

// Input: N = 4, K = 5
// Output: 1
// Explanation:
// row 1: 0
// row 2: 01
// row 3: 0110
// row 4: 01101001

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K%2 == 0 {
		if kthGrammar(N-1, K/2) == 0 {
			return 1
		}
		return 0
	}
	return kthGrammar(N-1, (K+1)/2)
}

func KthGrammarSolution() {
	log.Println("con meo", kthGrammar(5, 10))
}

// UNIQUE BINARY SEARCH TREES II
// Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

// Example:
// Input: 3
// Output:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// Explanation:
// The above output corresponds to the 5 unique BST's shown below:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

func generateTrees(n int) []*TreeNode {
	list := []*TreeNode{}
	if n == 0 {
		return list
	}
	list = generateRecursive(1, n)
	log.Println(list)
	return list
}

func generateRecursive(start int, end int) []*TreeNode {
	list := []*TreeNode{}
	if start > end {
		return list
	}
	for i := start; i <= end; i++ {
		leftSubTree := generateRecursive(start, i-1)
		rightSubTree := generateRecursive(i+1, end)

		for j := 0; j < len(leftSubTree); j++ {
			left := leftSubTree[j]
			for k := 0; k < len(rightSubTree); k++ {
				right := rightSubTree[k]
				node := TreeNode{i, left, right}
				list = append(list, &node)
			}
		}
	}
	return list
}
