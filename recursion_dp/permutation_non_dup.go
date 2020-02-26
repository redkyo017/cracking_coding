package recursion_dp

func GetPermutations(s string) *[]string {
	permutations := []string{}
	if len(s) == 0 {
		return &permutations
	}
	first := s[0]
	remainder := s[1:]
	words := GetPermutations(string(remainder))

	if len(*words) == 0 {
		*words = append(*words, string(first))
		return words
	}
	for _, word := range *words {
		for i := 0; i <= len(word); i++ {
			head := word[:i]
			tail := word[i:]
			newStr := string(head) + string(first) + string(tail)
			permutations = append(permutations, newStr)
		}
	}
	return &permutations
}
