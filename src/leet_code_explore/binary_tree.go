package leet_code_explore

import "log"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func preorderTraversal(root *TreeNode1) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	// visited := map[*TreeNode1]bool{}
	// stack := []*TreeNode1{}
	current := root
	for current != nil {
		result = append(result, current.Val)
	}
	return result
}

func PreorderImplement() {
	tn1 := TreeNode1{1, nil, nil}
	tn2 := TreeNode1{2, nil, nil}
	tn3 := TreeNode1{3, nil, nil}
	tn1.Right = &tn2
	tn2.Left = &tn3
	log.Println(preorderTraversal(&tn1))
}
