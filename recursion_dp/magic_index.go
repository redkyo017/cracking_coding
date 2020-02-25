package recursion_dp

import (
	"math"
)

func MagicIndex(a []int) int {
	return magicFastNotUnique(a, 0, len(a)-1)
}

func magicSlow(a []int) int {
	for i, v := range a {
		if i == v {
			return i
		}
	}
	return -1
}

func magicFastUnique(a []int, start, end int) int {
	if end < start {
		return -1
	}

	mid := (start + end) / 2
	if a[mid] == mid {
		return mid
	} else if a[mid] > mid {
		return magicFastUnique(a, start, mid-1)
	} else {
		return magicFastUnique(a, mid+1, end)
	}
}

func magicFastNotUnique(a []int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	midValue := a[mid]

	if mid == midValue {
		return mid
	}
	leftIndex := math.Min(float64(mid-1), float64(midValue))
	left := magicFastNotUnique(a, start, int(leftIndex))
	if left >= 0 {
		return left
	}
	rightIndex := math.Max(float64(mid+1), float64(midValue))
	right := magicFastNotUnique(a, int(rightIndex), end)
	return right
}
