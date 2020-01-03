package arraystrings

func IsStringRotation(s1 string, s2 string) bool {
	l := len(s1)
	if l == len(s2) && l > 0 {
		s1s1 := s1 + s1
		return isSubstring(s1s1, s2)
	}
	return false
}

// check if s2 is substring of s1
func isSubstring(s1 string, s2 string) bool {
	m := len(s2)
	n := len(s1)

	for i := 0; i < n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if s1[i+j] != s2[j] {
				break
			}
		}
		if j == m {
			return true
		}
	}
	return false
}
