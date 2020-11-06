package recursion_dp

func GetPermutationWithDuplicates(s string) *[]string {
	permutations := []string{}
	if len(s) == 0 {
		return &permutations
	}
	hashMap := map[string]int{}
	for _, c := range s {
		if _, ok := hashMap[string(c)]; ok {
			hashMap[string(c)]++
		} else {
			hashMap[string(c)] = 1
		}
	}
	GetPerWithDup(hashMap, "", len(s), &permutations)
	return &permutations
}

func GetPerWithDup(m map[string]int, prefix string, remain int, result *[]string) {
	if remain == 0 {
		*result = append(*result, prefix)
		return
	}
	for c, count := range m {
		if count > 0 {
			m[c]--
			GetPerWithDup(m, prefix+c, remain-1, result)
			m[c]++
		}
	}
}
