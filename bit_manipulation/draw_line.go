package bit_manipulation

import "log"

func DrawLine(screen []byte, width int, x1 int, x2 int, y int) {
	startOffset := x1 % 8
	first_full_byte := x1 / 8

	if startOffset != 0 {
		first_full_byte++
	}

	endOffset := x2 % 8
	last_full_byte := x2 / 8

	if endOffset != 7 {
		last_full_byte--
	}

	log.Println("con co", startOffset, first_full_byte)
}
