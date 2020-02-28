package recursion_dp

import "log"

func GenerateParenthese(n int) []string {
	result := FastGen(n)
	return result
}

func SlowGens(remain int) []string {
	parentheses := []string{}
	if remain == 0 {
		parentheses = append(parentheses, "")
	} else {
		prev := SlowGens(remain - 1)
		for _, str := range prev {
			for i := 0; i < len(str); i++ {
				if string(str[i]) == "(" {
					left := str[:i+1]
					right := str[i+1:]
					parentheses = append(parentheses, string(left)+"()"+string(right))
				}
			}
			parentheses = append(parentheses, "()"+str)
		}
	}

	return parentheses
}

func FastGen(remain int) []string {
	parentheses := []string{}
	char := make([]byte, remain*2)
	AddParents(&parentheses, remain, remain, char, 0)
	return parentheses
}

func AddParents(parentheses *[]string, leftRemain, rightRemain int, char []byte, index int) {
	if leftRemain < 0 || rightRemain < leftRemain {
		return
	}
	if leftRemain == 0 && rightRemain == 0 {
		log.Println("con co", string(char))
		*parentheses = append(*parentheses, string(char))
	} else {
		log.Println("con co index", index)
		char[index] = []byte("(")[0]
		AddParents(parentheses, leftRemain-1, rightRemain, char, index+1)
		log.Println("con co left", string(char))
		char[index] = []byte(")")[0]
		AddParents(parentheses, leftRemain, rightRemain-1, char, index+1)
		log.Println("con co right", string(char))
	}
}
