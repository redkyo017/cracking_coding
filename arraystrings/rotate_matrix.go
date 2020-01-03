package arraystrings

import "log"

func RotateMatrix(matrix [][]int) ([][]int, bool) {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return matrix, false
	}
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		start := layer
		end := n - 1 - layer

		for i := start; i < end; i++ {
			offset := i - start

			top := matrix[start][i]
			// left -> top
			matrix[start][i] = matrix[end-offset][start]
			// bottom -> left
			matrix[end-offset][start] = matrix[end][end-offset]
			// right -> bottom
			matrix[end][end-offset] = matrix[i][end]
			// bottom -> top
			matrix[i][end] = top
		}
	}
	for _, v := range matrix {
		log.Println("con co:", v)
	}
	return matrix, true
}
