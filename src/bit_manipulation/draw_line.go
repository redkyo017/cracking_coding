package bit_manipulation

import "log"

func DrawLine(screen []byte, width int, x1 int, x2 int, y int) {

	for i := 0; i < len(screen); i += width {
		j := i + width
		if j < len(screen) {
			log.Printf("%08b", screen[i:j])
		} else {
			log.Printf("%08b", screen[i:])
		}
	}
	log.Println("--------------------------")

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
	heightPos := (width / 8) * y
	log.Println("height pos", heightPos)
	log.Println("con co be be", startOffset, endOffset, first_full_byte, last_full_byte)
	for b := first_full_byte; b <= last_full_byte; b++ {
		log.Printf("con heo: %d : %08b \n", screen[(width/8)*y+b], screen[(width/8)*y+b])
		screen[(width/8)*y+b] = 0xFF
	}
	log.Printf("first full %d : %08b\n", first_full_byte, first_full_byte)
	log.Printf("last full %d : %08b\n", last_full_byte, last_full_byte)

	startMask := 0xFF >> uint(startOffset)
	endMask := ^(0xFF >> uint(endOffset+1))

	log.Printf("startMask %d :  %08b\n", startMask, startMask)
	log.Printf("endMask %d : %08b\n", endMask, endMask)
	log.Println("con meo", (x1 / 8), (x2 / 8), (x1/8) == (x2/8))
	if (x1 / 8) == (x2 / 8) {
		mask := startMask & endMask
		log.Printf("mask if x1 == x2 :%08b\n", mask)
		screen[(width/8)*y+(x1/8)] |= byte(mask)
	} else {
		if startOffset != 0 {
			byteNumber := (width/8)*y + first_full_byte - 1
			screen[byteNumber] |= byte(startMask)
		}
		if endOffset != 7 {
			byteNumber := (width/8)*y + last_full_byte + 1
			screen[byteNumber] |= byte(endMask)
		}
	}

	for i := 0; i < len(screen); i += width {
		j := i + width
		if j < len(screen) {
			log.Printf("%08b", screen[i:j])
		} else {
			log.Printf("%08b", screen[i:])
		}
	}
}
