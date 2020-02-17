package bit_manipulation

import (
	"log"
	"strconv"
)

func UpdateBits(n, m, i, j int64) int64 {
	allOnes := 0xFF
	left := allOnes << uint(j+1)
	log.Printf("left %08b", left)
	right := 1<<uint(i) - 1 // have 1s after position i
	log.Printf("right %08b", right)
	mask := int64(left | right)
	log.Printf("mask %08b", mask)

	// can clear the bits now after having the mask.
	nCleared := n & mask
	log.Printf("nCleared %08b\n", nCleared)

	// shift m so that it lines up with bits j through i.
	mShifted := m << uint(i)
	log.Printf("mShifted %08b\n", mShifted)
	// merge them together and return.
	result := nCleared | mShifted
	log.Printf("con co: %b\n", result)
	return result
}

func BinToInt(b string) int64 {
	i, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return -1
	}
	return i
}
