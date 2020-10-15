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
	listSteps := make([]int, len(nums))
	for i := 0; i<len(nums); i++ {
		listSteps[i] = 0
	}
	for i := 0; i<len(nums); i++ {
		stepAtIndex := nums[i]
		if stepAtIndex == 0 {
			continue
		}
		for step := 1; step <= stepAtIndex; step++ {
			if (step+i) >= len(nums) {
				break
			}
			val := listSteps[i+step]
			if step > val {
				listSteps[i+step] = step
			}
		}
	}
	minStep := 0
	index := len(listSteps)-1
	for index > 0 {
		index = index - listSteps[index]
		minStep++
	}
	log.Println("con meo", listSteps, minStep)
	return minStep
}