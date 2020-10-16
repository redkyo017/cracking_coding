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

import (
	"log"
	// "math/rand"
	// "time"
)

func JumpII() {
	nums := []int{5,6,5,3,9,8,3,1,2,8,2,4,8,3,9,1,0,9,4,6,5,9,8,7,4,2,1,0,2}
	// nums := []int{3,2,4,1,1,1,1}
	// rand.Seed(time.Now().Unix())
	// nums := rand.Perm(100)
	// nums := []int{3,2,1}
	// nums := []int{2,3,1}
	// log.Println("con co", nums)
	log.Println(Solution2(nums))
}


func Solution1(nums []int) int {
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
	// log.Println("con meo", listSteps, minStep)
	return minStep
}

func Solution2(nums []int) int {
	if len(nums) <= 1 { 
		return 0
	}
	minStep := 0
	index := 0
	for index < len(nums)-1 {
		stepAtIndex := nums[index]
		if stepAtIndex == 0 {
			index++
			continue
		}
		atIndex := index
		maxCanJupm := 0
		for i := 1;i <= stepAtIndex; i++ {
			nextStep := i+atIndex
			if nextStep >= len(nums)-1 {
				index = nextStep
				continue
				// break
			}
			if nums[nextStep] + nextStep >= maxCanJupm {
				maxCanJupm = nums[nextStep] + nextStep
				index = nextStep
			}
		}
		minStep++
	}
	// log.Println("con co", minStep, index)
	return minStep
}