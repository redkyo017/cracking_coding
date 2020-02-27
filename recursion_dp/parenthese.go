package recursion_dp

func GenerateParenthese(n int) []string {
	result := SlowGens(n)
	return *result
}

func SlowGens(remain int) *[]string {
	parentheses := []string{}

	if remain < 0 {
		parentheses = append(parentheses, "")
	} else {
		prev := SlowGens(remain - 1)
		for _, str := range *prev {
			for i := 0; i < len(str); i++ {
				if string(str[i]) == "(" {
					left := str[:i]
					right := str[i:]

					parentheses = append(parentheses, string(left)+"()"+string(right))
				}
			}
			parentheses = append(parentheses, "()"+str)
		}
	}

	return &parentheses
}
