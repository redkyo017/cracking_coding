package recursion_dp

func CountWays1(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return CountWays1(n-1) + CountWays1(n-2) + CountWays1(n-3)
	}
}

func CountWays2(n int) int {
	memo := map[int]int{}
	return CountWaysWithMemorization(n, memo)
}

func CountWaysWithMemorization(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if v, ok := memo[n]; ok {
		return v
	}

	memo[n] = CountWaysWithMemorization(n-1, memo) + CountWaysWithMemorization(n-2, memo) + CountWaysWithMemorization(n-3, memo)

	return memo[n]
}
