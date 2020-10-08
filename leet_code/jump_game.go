// tag: HARD
// Given an array of non-negative integers, you are initially positioned at the first index of the array.

// Each element in the array represents your maximum jump length at that position.

// Your goal is to reach the last index in the minimum number of jumps.

// Example:

// Input: [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2.
//     Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Note:

// You can assume that you can always reach the last index.
package leet_code

import "log"

// func main() {
// 	arr := []int{2, 3, 1, 1, 4}
// 	log.Println("con co be be", jump(arr))
// }


func JumpII(nums []int) int {
	if len(nums) <= 1 { 
		return 0
	}
	list_step := make(map[int][][]int, len(nums))
	list_step[0] = append([][]int{}, []int{0})
	list_step[1] = append([][]int{}, []int{1})
	
	for i := 2; i < len(nums);i++ {
		stepsAtIndex := [][]int{}
		// log.Println("con heo", i, nums[i], stepsAtIndex)
		for j := i-1; j >= 0; j-- {
			if nums[j] >= i-j {
				// log.Println("can step", j)
				stepAtJ := list_step[j]
				for _,v := range stepAtJ {
					newStep := append(v, i-j)
					stepsAtIndex = append(stepsAtIndex, newStep)
				}
			}
		}
		list_step[i] = stepsAtIndex
	}
	stepAtEnd := list_step[len(nums)-1]
	const MaxUint = ^uint(0)
	minStep := int(MaxUint >> 1) 
	for _,v := range stepAtEnd {
		if len(v) < minStep {
			minStep = len(v)
		}
	}
	// log.Println("con co", minStep)
	return minStep
}
