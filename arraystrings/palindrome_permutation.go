package arraystrings

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

	// solution 2: use array int to count
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

	// bit = 1011 && 0100 == 0 => bit = 1011 |= 0100 = 1111
	// bit = 1011 && 1000 == 1000 => bit = 1011 &= ~1000(0111) = 0011
	bitVector := 0
	for _, v := range str {
		mask := 1 << uint(v)
		if bitVector&mask == 0 {
			bitVector |= mask
		} else {
			bitVector &= ^mask
		}
	}
	return bitVector == 0 || (bitVector&(bitVector-1) == 0)
	// return true
}
