package arraystrings

import (
	"log"
)

func IsPermutationOfPalindrome(str string) bool {
	// solution 1: hashmap to count max one odd characters
	// frequentMaps := map[string]int{}
	// for _, v := range str {
	// 	frequentMaps[string(v)]++
	// }
	// maxOneOdd := false
	// for _, v := range frequentMaps {
	// 	if v%2 == 1 {
	// 		if maxOneOdd {
	// 			return false
	// 		}
	// 		maxOneOdd = true
	// 	}
	// }
	// return true

	// // solution 2: use array int to count
	// countOdd := 0
	// countTable := make([]int, 255)
	// for _, v := range str {
	// 	countTable[int(v)]++
	// 	if countTable[int(v)]%2 == 1 {
	// 		countOdd++
	// 	} else {
	// 		countOdd--
	// 	}
	// }
	// return countOdd <= 1

	// solution 3: bitwise
	bitVector := 0
	var aChar uint = 97
	for _, v := range str {
		index := uint(v) - aChar
		if index < 0 {
			break
		}

		mask := 1 << index
		if bitVector&mask == 0 {
			bitVector |= mask
		} else {
			bitVector &= ^mask
		}
		log.Printf("con co %b", bitVector)
	}
	log.Println(bitVector == 0, (bitVector&(bitVector-1) == 0))
	return bitVector == 0 || (bitVector&(bitVector-1) == 0)
	// return true
}
