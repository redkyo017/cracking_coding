package leet_code_explore

import (
	"log"
	"math"
)

// DEVIDE AND CONQUER
// MERGE SORT
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := len(nums) / 2
	left_list := sortArray(nums[:pivot])
	right_list := sortArray(nums[pivot:])
	return mergeList(left_list, right_list)
}
func mergeList(left []int, right []int) []int {
	left_cursor, right_cursor := 0, 0
	res := []int{}
	for left_cursor < len(left) && right_cursor < len(right) {
		if left[left_cursor] < right[right_cursor] {
			res = append(res, left[left_cursor])
			left_cursor++
		} else {
			res = append(res, right[right_cursor])
			right_cursor++
		}
	}
	res = append(res, left[left_cursor:]...)
	res = append(res, right[right_cursor:]...)
	return res
}

func SortArrayImplement() {
	arr := []int{1, 5, 3, 2, 8, 7, 6, 4}
	log.Println("before sort", arr)
	log.Println("after sort", sortArray(arr))
}

// BINARY SEARCH TREE
const MININT = math.MinInt32
const MAXINT = math.MaxInt32

func isValidBST(root *TreeNode) bool {
	log.Println("con meo", MININT, MAXINT)
	return isValidBSTUtil(root, MININT, MAXINT)
}

func isValidBSTUtil(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if (node.Val < min) || (node.Val > max) {
		return false
	}
	return isValidBSTUtil(node.Left, min, node.Val-1) && isValidBSTUtil(node.Right, node.Val+1, max)
}

func IsValidBSTImplement() {
	// node1 := TreeNode{1, nil, nil}
	// node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node4 := TreeNode{4, nil, nil}
	node5 := TreeNode{5, nil, nil}
	node6 := TreeNode{6, nil, nil}
	node7 := TreeNode{7, nil, nil}

	node5.Left = &node4
	node5.Right = &node6
	node6.Right = &node7
	node6.Left = &node3
	log.Println("con meo", isValidBST(&node5))
}

// SEARCH A 2D MATRIX II
// Example 1:
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// Output: true
// Example 2:
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// Output: false
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row <= 0 {
		return false
	}
	col := len(matrix[0])
	if col <= 0 {
		return false
	}
	colPivot, rowPivot := col/2, row/2
	if matrix[rowPivot][colPivot] == target {
		return true
	} else if matrix[rowPivot][colPivot] < target {
		return searchMatrix(matrix[:rowPivot][:colPivot], target) || searchMatrix(matrix[rowPivot:][:colPivot], target) || searchMatrix(matrix[:rowPivot][colPivot:], target)
	} else {
		return searchMatrix(matrix[rowPivot:][colPivot:], target) || searchMatrix(matrix[rowPivot:][:colPivot], target) || searchMatrix(matrix[:rowPivot][colPivot:], target)
	}
	// return false
}

// func SearchMatrixUtil(matrix [][]int, row int, col int, target int) bool {
// 	return false
// }

func SearchMatrixImplement() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	log.Println("search", searchMatrix(matrix, 5))
}
