package recursion_dp

func GetSubsets(set []int, index int) *[][]int {
	if index == len(set) {
		return &[][]int{}
	}
	allSubsets := GetSubsets(set, index+1)
	item := set[index]
	moreSubset := [][]int{}

	if len(*allSubsets) == 0 {
		*allSubsets = append(*allSubsets, []int{item})
	} else {
		for _, subset := range *allSubsets {
			// newSubset := []int{}
			subset = append(subset, item)
			moreSubset = append(moreSubset, subset)
		}
	}
	*allSubsets = append(*allSubsets, moreSubset...)
	return allSubsets
}
