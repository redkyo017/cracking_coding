package leet_code_explore

import (
	"fmt"
	"log"
	"math"
	"strconv"
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
	rowPivot, colPivot := row/2, col/2
	if rowPivot == 0 && colPivot == 0 && matrix[rowPivot][colPivot] != target {
		return false
	}
	top := matrix[:rowPivot]
	bottom := matrix[rowPivot:]
	if matrix[rowPivot][colPivot] == target {
		return true
	}
	topRight, bottomLeft := [][]int{}, [][]int{}
	if target < matrix[rowPivot][colPivot] {
		topLeft := [][]int{}
		for _, t := range top {
			topLeft = append(topLeft, t[:colPivot])
			topRight = append(topRight, t[colPivot:])
		}
		for _, b := range bottom {
			bottomLeft = append(bottomLeft, b[:colPivot])
		}
		return searchMatrix(topLeft, target) || searchMatrix(topRight, target) || searchMatrix(bottomLeft, target)
	}
	if target > matrix[rowPivot][colPivot] {
		bottomRight := [][]int{}
		for _, t := range top {
			topRight = append(topRight, t[colPivot:])
		}
		for _, b := range bottom {
			bottomLeft = append(bottomLeft, b[:colPivot])
			bottomRight = append(bottomRight, b[colPivot:])
		}
		return searchMatrix(bottomRight, target) || searchMatrix(topRight, target) || searchMatrix(bottomLeft, target)
	}
	return false
}

func SearchMatrixImplement() {
	// matrix := [][]int{
	// 	{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, 6, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30},
	// }
	// matrix := [][]int{
	// 	{-1, 3},
	// }
	matrix := [][]int{
		{5},
		{6},
	}
	log.Println("search", searchMatrix(matrix, 5))
}

// BACK TRACKING TEMPLATE
// def backtrack(candidate):
//     if find_solution(candidate):
//         output(candidate)
//         return

//     # iterate all possible candidates.
//     for next_candidate in list_of_candidates:
//         if is_valid(next_candidate):
//             # try this partial candidate solution
//             place(next_candidate)
//             # given the candidate, explore further.
//             backtrack(next_candidate)
//             # backtrack
//             remove(next_candidate)

func totalNQueens(n int) int {
	chessTable := [][]bool{}
	for i := 0; i < n; i++ {
		row := []bool{}
		for j := 0; j < n; j++ {
			row = append(row, false)
		}
		chessTable = append(chessTable, row)
	}
	var count int
	backtrackQueen(&chessTable, 0, &count)
	return count
}

func isValidCell(table *[][]bool, row int, col int) bool {
	if (*table)[row][col] == true {
		return false
	}
	n := len(*table)
	for i := 0; i < n; i++ {
		// row & col
		if ((*table)[i][col] == true) || ((*table)[row][i] == true) {
			return false
		}
	}
	// diagonal
	for i := 0; i <= col; i++ {
		top := row - i
		bot := row + i
		prevCol := col - i
		if (top >= 0 && (*table)[top][prevCol] == true) || (bot < n && (*table)[bot][prevCol] == true) {
			return false
		}
	}
	return true
}

func backtrackQueen(table *[][]bool, col int, count *int) {
	n := len(*table)
	if col == n {
		// log.Println("con co be be", table)
		(*count)++
		return
	}
	for i := 0; i < n; i++ {
		if isValidCell(table, i, col) {
			(*table)[i][col] = true
			backtrackQueen(table, col+1, count)
			(*table)[i][col] = false
		}
	}
}
func TotalQueenImplement() {
	n := 8
	log.Println("total queens", totalNQueens(n))
}

// Input: board = [
// 	["5","3",".",".","7",".",".",".","."],
// 	["6",".",".","1","9","5",".",".","."],
// 	[".","9","8",".",".",".",".","6","."],
// 	["8",".",".",".","6",".",".",".","3"],
// 	["4",".",".","8",".","3",".",".","1"],
// 	["7",".",".",".","2",".",".",".","6"],
// 	[".","6",".",".",".",".","2","8","."],
// 	[".",".",".","4","1","9",".",".","5"],
// 	[".",".",".",".","8",".",".","7","9"]
// ]
// Output: [
// 	["5","3","4","6","7","8","9","1","2"],
// 	["6","7","2","1","9","5","3","4","8"],
// 	["1","9","8","3","4","2","5","6","7"],
// 	["8","5","9","7","6","1","4","2","3"],
// 	["4","2","6","8","5","3","7","9","1"],
// 	["7","1","3","9","2","4","8","5","6"],
// 	["9","6","1","5","3","7","2","8","4"],
// 	["2","8","7","4","1","9","6","3","5"],
// 	["3","4","5","2","8","6","1","7","9"]
// ]
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit or '.'

func solveSudoku(board [][]byte) {
	backtrackSudoku(&board, 0, 0)
}
func intToByte(n int) byte {
	numChar := fmt.Sprintf("%d", n)
	return []byte(numChar)[0]
}
func byteToInt(c byte) int {
	num, _ := strconv.Atoi(string(c))
	return num
}
func isValid(board *[][]byte, row int, col int, num int) bool {
	numByte := intToByte(num)
	for i := 0; i < 9; i++ {
		if (*board)[row][i] == numByte || (*board)[i][col] == numByte {
			return false
		}
	}
	startRow := row - (row % 3)
	startCol := col - (col % 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[i+startRow][j+startCol] == numByte {
				return false
			}
		}
	}
	return true
}
func backtrackSudoku(board *[][]byte, row int, col int) bool {
	if col == 9 && row == 8 {
		return true
	}
	// consider next candidate
	if col == 9 {
		col = 0
		row++
	}

	if (*board)[row][col] != '.' {
		return backtrackSudoku(board, row, col+1)
	}

	for i := 1; i <= 9; i++ {
		if isValid(board, row, col, i) {
			(*board)[row][col] = intToByte(i)
			if backtrackSudoku(board, row, col+1) {
				return true
			}
		}
		(*board)[row][col] = '.'
	}
	return false
}

func SolveSudokuImplement() {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),

		// []byte("..9748..."),
		// []byte("7........"),
		// []byte(".2.1.9..."),
		// []byte("..7...24."),
		// []byte(".64.1.59."),
		// []byte(".98...3.."),
		// []byte("...8.3.2."),
		// []byte("........6"),
		// []byte("...2759.."),
	}
	solveSudoku(board)
	for _, rows := range board {
		log.Println(string(rows))
	}
}

// [
// 	[".",".","9","7","4","8",".",".","."],
// 	["7",".",".",".",".",".",".",".","."],
// 	[".","2",".","1",".","9",".",".","."],
// 	[".",".","7",".",".",".","2","4","."],
// 	[".","6","4",".","1",".","5","9","."],
// 	[".","9","8",".",".",".","3",".","."],
// 	[".",".",".","8",".","3",".","2","."],
// 	[".",".",".",".",".",".",".",".","6"],
// 	[".",".",".","2","7","5","9",".","."]
// ]

// [
// 	["3","1","9","7","4","8","6","5","2"],
// 	["7","4","3","6","5","2","1","8","9"],
// 	["6","2","5","1","3","9","8","7","4"],
// 	["5","3","7","9","8","6","2","4","1"],
// 	["2","6","4","3","1","7","5","9","8"],
// 	["1","9","8","5","2","4","3","6","7"],
// 	["9","7","1","8","6","3","4","2","5"],
// 	["8","5","2","4","9","1","7","3","6"],
// 	["4","8","6","2","7","5","9","1","3"]
// ]

// [
// 	["5","1","9","7","4","8","6","3","2"],
// 	["7","8","3","6","5","2","4","1","9"],
// 	["4","2","6","1","3","9","8","7","5"],
// 	["3","5","7","9","8","6","2","4","1"],
// 	["2","6","4","3","1","7","5","9","8"],
// 	["1","9","8","5","2","4","3","6","7"],
// 	["9","7","5","8","6","3","1","2","4"],
// 	["8","3","2","4","9","1","7","5","6"],
// 	["6","4","1","2","7","5","9","8","3"]
// ]

func CombineImplement() {
	n, k := 4, 2
	res := combine(n, k)
	// for _, v := range res {
	log.Println(res)
	// }
}

func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrackCombine(&res, []int{}, k, n, 1)
	return res
}

func backtrackCombine(combinations *[][]int, combination []int, combinationLen int, numLen int, currentNum int) {
	if len(combination) == combinationLen {
		(*combinations) = append((*combinations), combination)
		return
	}
	for i := currentNum; i <= numLen; i++ {
		var combs []int
		combs = append(combs, combination...)
		combs = append(combs, i)
		combination = append(combination, i)
		backtrackCombine(combinations, combination, combinationLen, numLen, i+1)
	}
}
