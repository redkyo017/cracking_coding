package recursion_dp

import (
	"fmt"
	"log"
)

func stringToBool(c string) bool {
	if c == "1" {
		return true
	}
	return false
}

func CountEval(s string, eval bool, hmap map[string]int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if stringToBool(s) == eval {
			return 1
		}
		return 0
	}
	key := fmt.Sprintf("%s_%t", s, eval)
	if v, ok := hmap[key]; ok {
		return v
	}

	ways := 0
	for i := 1; i < len(s); i += 2 {
		char := string(s[i])

		left := string(s[:i])
		right := string(s[i+1:])

		leftTrue := CountEval(left, true, hmap)
		leftFalse := CountEval(left, false, hmap)

		rightTrue := CountEval(right, true, hmap)
		rightFalse := CountEval(right, false, hmap)

		total := (leftTrue + leftFalse) * (rightTrue + rightFalse)
		totalTrue := 0
		if char == "^" {
			totalTrue = leftTrue*rightFalse + leftFalse*rightTrue
		} else if char == "&" {
			totalTrue = leftTrue + rightTrue
		} else if char == "|" {
			totalTrue = leftTrue*rightTrue + leftFalse*rightTrue + leftTrue*rightFalse
		}
		subWays := 0
		if eval {
			subWays = totalTrue
		} else {
			subWays = total - totalTrue
		}
		ways += subWays
	}
	hmap[key] = ways
	return ways
}

func ImplementCountEval() {
	s := "1^0|0|1"
	result := false
	hmap := make(map[string]int)

	log.Println("count Eval", CountEval(s, result, hmap))
}
