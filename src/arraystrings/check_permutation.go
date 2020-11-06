package arraystrings

import (
	"sort"
)

// "sort"

func CheckPermutation(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	// cach 1 : sort 2 string xong so sanh
	b1 := []byte(str1)
	b2 := []byte(str2)

	sort.Slice(b1, func(i int, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i int, j int) bool { return b2[i] < b2[j] })
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	// cach 2: check identical character count
	// letters := [128]int{}
	// for _, v := range str1 {
	// 	letters[int(v)]++
	// }
	// for i := 0; i < len(str2); i++ {
	// 	c := int(str2[i])
	// 	letters[c]--
	// 	if letters[c] < 0 {
	// 		return false
	// 	}
	// }
	return true
}
