package main

import (
	"cracking_coding/arraystrings"
	"log"
)

func main() {
	log.Println(arraystrings.SetZeroMatrix([][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 0, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}))
}
