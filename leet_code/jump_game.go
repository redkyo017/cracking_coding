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

func JumpII(arr []int) int {
	list_step := make(map[int][][]int, len(arr))
	list_step[0] = append([][]int{}, []int{0})
	list_step[1] = append([][]int{}, []int{1})
	
	for i := 2; i < len(arr);i++ {
		stepsAtIndex := [][]int{}
		log.Println("con heo", i, arr[i], stepsAtIndex)
	}
	log.Println("con co", list_step, len(list_step))
	return 0
}
