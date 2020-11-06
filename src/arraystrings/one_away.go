package arraystrings

import (
	"log"
	"math"
)

// There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.

func OneEditAway(s1 string, s2 string) bool {
	log.Println(s1, s2)
	// solution 1:
	// if len(s1) == len(s2) {
	// 	return oneReplaceAway(s1, s2)
	// } else if len(s1)+1 == len(edit) {
	// 	return oneInsertAway(s2, s1)
	// } else if len(s1)-1 == len(s2) {
	// 	return oneInsertAway(s1, s2)
	// }
	// return false

	// solution 2:
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}
	st1, st2 := "", ""
	if len(s1) < len(s2) {
		st1 = s1
		st2 = s2
	} else {
		st1 = s2
		st2 = s1
	}

	foundDifferent := false
	index1, index2 := 0, 0
	for index1 < len(st1) && index2 < len(st2) {
		if (st1[index1]) != st2[index2] {
			if foundDifferent {
				return false
			}
			foundDifferent = true

			if len(s1) == len(s2) {
				index1++
			}
		} else {
			index1++
		}
		index2++
	}
	// solution 2.1
	// foundDifferent := false
	// index1, index2 := 0, 0
	// for index1 < len(s1) && index2 < len(s2) {
	// 	if s1[index1] != s2[index2] {
	// 		if foundDifferent {
	// 			return false
	// 		}
	// 		foundDifferent = true
	// 		if len(s1) > len(s2) {
	// 			index1++
	// 		} else if len(s1) < len(s2) {
	// 			index2++
	// 		} else {
	// 			index1++
	// 			index2++
	// 		}
	// 	} else {
	// 		index1++
	// 		index2++
	// 	}
	// }
	return true
}

func oneReplaceAway(s1 string, s2 string) bool {
	foundDifferent := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if foundDifferent {
				return false
			}
			foundDifferent = true
		}
	}
	return true
}

func oneInsertAway(s1 string, s2 string) bool {
	index1, index2 := 0, 0
	for index1 < len(s1) && index2 < len(s2) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}
