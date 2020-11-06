package arraystrings

import "log"

func SetZeroMatrix(matrix [][]int) bool {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for row := 0; row < len(rows); row++ {
		if rows[row] {
			for col := 0; col < len(matrix[0]); col++ {
				matrix[row][col] = 0
			}
		}
	}
	for col := 0; col < len(cols); col++ {
		if cols[col] {
			for row := 0; row < len(matrix); row++ {
				matrix[row][col] = 0
			}
		}
	}
	for _, v := range matrix {
		log.Println("con meo", v)
	}
	return false
}
