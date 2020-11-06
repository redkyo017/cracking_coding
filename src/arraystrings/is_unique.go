package arraystrings

import "log"

func IsUniqueString(str string) bool {
	log.Println("check if is unique solution")
	if len(str) > 128 {
		return false
	}

	char_set := [128]bool{}
	for _, v := range str {
		if char_set[int(v)] {
			return false
		}
		char_set[int(v)] = true
	}
	log.Println(char_set)
	return true
}

func IsUniqueStringBitwise(str string) bool {
	log.Println("check if is unique solution by bitwise")
	var aChar uint = 97
	checker := 0
	for _, v := range str {
		val := uint(v) - aChar
		if checker&(1<<val) > 0 {
			return false
		}
		checker |= 1 << val
	}
	return true
}
