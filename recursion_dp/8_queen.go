package recursion_dp

import (
	"log"
	"math"
)

const GRID_SIZE = 8

func EightQueen() {
	results := [][]int{}
	columns := make([]int, GRID_SIZE)
	PlaceQueen(0, columns, &results)
	log.Println("con co", results)
}

func PlaceQueen(row int, columns []int, results *[][]int) {
	if row == GRID_SIZE {
		(*results) = append(*results, columns)
	} else {
		for col := 0; col < GRID_SIZE; col++ {
			if checkValid(columns, row, col) {
				columns[row] = col
				PlaceQueen(row+1, columns, results)
			}
		}
	}
}

func checkValid(columns []int, row1 int, col1 int) bool {
	for row2 := 0; row2 < row1; row2++ {
		column2 := columns[row2]

		if column2 == col1 {
			return false
		}

		columDistance := int(math.Abs(float64(column2) - float64(col1)))
		rowDistance := row1 - row2
		if columDistance == rowDistance {
			return false
		}
	}
	return true
}
