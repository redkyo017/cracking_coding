package leet_code_explore

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type StraightQueue struct {
	value []int
}

func (q *StraightQueue) StraightEnqueue(value int) bool {
	q.value = append(q.value, value)
	return true
}

func (q *StraightQueue) StraightDequeue() bool {
	if q.IsEmpty() {
		return false
	}
	log.Println("dequeue value", q.value[0])
	q.value = q.value[1:]
	return true
}

func (q *StraightQueue) StraightFront() int {
	if q.IsEmpty() {
		return 0
	}
	return q.value[0]
}

func (q *StraightQueue) StraightRear() int {
	if q.IsEmpty() {
		return 0
	}
	return q.value[len(q.value)-1]
}

func (q *StraightQueue) IsEmpty() bool {
	return len(q.value) == 0
}

type CircularQueue struct {
	value [5]int
	head  int
	tail  int
}

func (q *CircularQueue) CircularEnqueue(value int) bool {
	if q.IsFull() {
		return false
	}
	if q.IsEmpty() {
		q.head = 0
	}
	q.tail = (q.tail + 1) % len(q.value)
	q.value[q.tail] = value
	return true
}

func (q *CircularQueue) CircularDequeue() bool {
	if q.IsEmpty() {
		return false
	}
	if q.head == q.tail {
		q.head = -1
		q.tail = -1
		return true
	}
	q.head = (q.head + 1) % len(q.value)
	log.Println("dequeue value", q.value[0])
	return true
}

func (q *CircularQueue) CircularFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.value[q.head]
}

func (q *CircularQueue) CircularRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.value[q.tail]
}

func (q *CircularQueue) IsFull() bool {
	return (q.tail+1)%len(q.value) == q.head
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == -1
}

func QueueImplementation() {
	// strait queue
	// squeue := StraightQueue{
	// 	value: []int{},
	// }
	// squeue.StraightEnqueue(1)
	// squeue.StraightEnqueue(1)
	// squeue.StraightEnqueue(2)
	// squeue.StraightEnqueue(3)
	// squeue.StraightEnqueue(5)
	// log.Println("enqueued", squeue.value)
	// squeue.StraightDequeue()
	// squeue.StraightDequeue()
	// log.Println("dequeued", squeue.value)

	// circular queue
	cqueue := CircularQueue{
		value: [5]int{},
		head:  -1,
		tail:  -1,
	}
	cqueue.CircularEnqueue(1)
	cqueue.CircularEnqueue(1)
	cqueue.CircularEnqueue(2)
	cqueue.CircularEnqueue(3)
	cqueue.CircularEnqueue(5)
	log.Println("enqueued", cqueue.CircularFront(), cqueue.CircularRear(), cqueue)
	log.Println(cqueue.CircularEnqueue(8))
	cqueue.CircularDequeue()
	cqueue.CircularDequeue()
	log.Println("dequeued", cqueue.CircularFront(), cqueue.CircularRear(), cqueue)
}

// TEMPLATE BFS - Breadth first search
// int BFS(Node root, Node target) {
//     Queue<Node> queue;  // store all nodes which are waiting to be processed
//     int step = 0;       // number of steps neeeded from root to current node
//     // initialize
//     add root to queue;
//     // BFS
//     while (queue is not empty) {
//         step = step + 1;
//         // iterate the nodes which are already in the queue
//         int size = queue.size();
//         for (int i = 0; i < size; ++i) {
//             Node cur = the first node in queue;
//             return step if cur is target;
//             for (Node next : the neighbors of cur) {
//                 add next to queue;
//             }
//             remove the first node from queue;
//         }
//     }
//     return -1;          // there is no path from root to target
// }

// PRACTISE BFS
func numIsland(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	visitedNode := make(map[string]bool)
	type nodeItem struct {
		row int
		col int
	}
	landNode := []nodeItem{}
	step := 0
	for r, lines := range grid {
		for c, value := range lines {
			if string(value) == "1" {
				landNode = append(landNode, nodeItem{r, c})
			}
		}
	}

	for _, v := range landNode {
		if _, ok := visitedNode[fmt.Sprintf("%d-%d", v.row, v.col)]; ok {
			continue
		}
		step++
		nodeQueue := append([]nodeItem{}, v)
		for len(nodeQueue) > 0 {
			size := len(nodeQueue)
			for i := 0; i < size; i++ {
				node := nodeQueue[0]
				row := node.row
				col := node.col
				if col+1 < len(grid[0]) && string(grid[row][col+1]) == "1" {
					if _, ok := visitedNode[fmt.Sprintf("%d-%d", row, col+1)]; !ok {
						nodeQueue = append(nodeQueue, nodeItem{row, col + 1})
						visitedNode[fmt.Sprintf("%d-%d", row, col+1)] = true
					}
				}
				if col-1 >= 0 && string(grid[row][col-1]) == "1" {
					if _, ok := visitedNode[fmt.Sprintf("%d-%d", row, col-1)]; !ok {
						nodeQueue = append(nodeQueue, nodeItem{row, col - 1})
						visitedNode[fmt.Sprintf("%d-%d", row, col-1)] = true
					}
				}

				if row+1 < len(grid) && string(grid[row+1][col]) == "1" {
					if _, ok := visitedNode[fmt.Sprintf("%d-%d", row+1, col)]; !ok {
						nodeQueue = append(nodeQueue, nodeItem{row + 1, col})
						visitedNode[fmt.Sprintf("%d-%d", row+1, col)] = true
					}
				}

				if row-1 >= 0 && string(grid[row-1][col]) == "1" {
					if _, ok := visitedNode[fmt.Sprintf("%d-%d", row-1, col)]; !ok {
						nodeQueue = append(nodeQueue, nodeItem{row - 1, col})
						visitedNode[fmt.Sprintf("%d-%d", row-1, col)] = true
					}
				}
				nodeQueue = nodeQueue[1:]
			}
		}
	}
	log.Println("step", step)
	return step
}

func NumIslandSolution() {
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	// grid := [][]byte{
	// 	[]byte("11110"),
	// 	[]byte("11010"),
	// 	[]byte("11000"),
	// 	[]byte("00000"),
	// }
	// grid := [][]byte{
	// 	[]byte("111"),
	// 	[]byte("010"),
	// 	[]byte("111"),
	// }
	// grid := [][]byte{
	// 	[]byte("10111"),
	// 	[]byte("10101"),
	// 	[]byte("11101"),
	// }
	// numIsland(grid)
	numIslandDFS(grid)
}

// You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

// The lock initially starts at '0000', a string representing the state of the 4 wheels.

// You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

// Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

// Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// Output: 6
// Input: deadends = ["8888"], target = "0009"
// Output: 1
// Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
// Output: -1
// Input: deadends = ["0000"], target = "8888"
// Output: -1

func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool)
	for _, v := range deadends {
		deadendsMap[v] = true
	}
	step := -1
	queueLocks := append([]string{}, "0000")
	visited := map[string]bool{"0000": true}
	for len(queueLocks) > 0 {
		step++
		size := len(queueLocks)
		log.Println("con meo", size)
		for i := 0; i < size; i++ {
			lockCase := queueLocks[0]
			if lockCase == target {
				return step
			}
			if _, ok := deadendsMap[lockCase]; ok {
				queueLocks = queueLocks[1:]
				continue
			}
			for k, numStr := range lockCase {
				num, _ := strconv.Atoi(string(numStr))
				toward := (num + 1)
				backward := (num - 1)
				if toward > 9 {
					toward = 0
				}
				if backward < 0 {
					backward = 9
				}
				towardCase := fmt.Sprintf("%s%d%s", lockCase[:k], toward, lockCase[k+1:])
				if _, ok := visited[towardCase]; !ok {
					queueLocks = append(queueLocks, towardCase)
					visited[towardCase] = true
				}
				backwardCase := fmt.Sprintf("%s%d%s", lockCase[:k], backward, lockCase[k+1:])
				if _, ok := visited[backwardCase]; !ok {
					queueLocks = append(queueLocks, backwardCase)
					visited[backwardCase] = true
				}
			}
			queueLocks = queueLocks[1:]
			// log.Println("con co", queueLocks)
		}
	}
	return -1
}

func OpenLockSolution() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	// deadends := []string{"8888"}
	// target := "0009"
	// deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	// target := "8888"
	// deadends := []string{"0000"}
	// target := "8888"
	log.Println("openLock: ", openLock(deadends, target))
}

// PERFECT SQUARES
// Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

// Example 1:
// Input: n = 12
// Output: 3
// Explanation: 12 = 4 + 4 + 4.

// Example 2:
// Input: n = 13
// Output: 2
// Explanation: 13 = 4 + 9.

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	step := 0
	queueNums := [][]int{{0, n}}

	for len(queueNums) > 0 {
		step++
		size := len(queueNums)
		log.Println("con meo", queueNums, size, step)
		for i := 0; i < size; i++ {
			prevPerfectSuareNum := queueNums[0][0]
			remain := queueNums[0][1]
			for j := int(math.Sqrt(float64(prevPerfectSuareNum))); j*j <= remain; j++ {
				perfectSquare := j * j
				if remain-perfectSquare == 0 {
					return step
				}
				if perfectSquare == 0 || remain-perfectSquare < 0 {
					continue
				}
				queueNums = append(queueNums, []int{j * j, remain - (j * j)})
			}
			queueNums = queueNums[1:]
		}
		log.Println("con co", queueNums)
	}
	return -1
}

func NumSquaresSolution() {
	n := 12
	log.Println(numSquares(n))
}

type MinStack struct {
	values [][]int
}

func (this *MinStack) Push(v int) {
	if len(this.values) == 0 {
		this.values = append(this.values, []int{v, v})
		return
	}
	top := this.values[len(this.values)-1]
	min := top[1]
	if v <= min {
		this.values = append(this.values, []int{v, v})
	} else {
		this.values = append(this.values, []int{v, min})
	}
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
}

func (this *MinStack) Top() int {
	if len(this.values) == 0 {
		return -1
	}
	top := this.values[len(this.values)-1]
	return top[0]
}

func (this *MinStack) GetMin() int {
	top := this.values[len(this.values)-1]
	return top[1]
}

func ImplementMinStack() {
	minStack := MinStack{
		values: [][]int{},
	}
	log.Println("con co be be", minStack)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	log.Println("min1", minStack.GetMin())
	minStack.Pop()
	log.Println("con co be be", minStack)
	log.Println("top", minStack.Top())
	log.Println("min2", minStack.GetMin())
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "([)]"
// Output: false

// Example 5:
// Input: s = "{[]}"
// Output: true

func isValidParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := []string{}
	mapPair := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for _, v := range s {
		parentheses := string(v)
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if mapPair[top] == parentheses {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, string(v))
	}
	return len(stack) == 0
}

func IsValidParenthesesSolution() {
	s := "{()}"
	log.Println("valid or not", isValidParentheses(s))
}

//   DAILY TEMPERATURES
// Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

// For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

// Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

func dailyTemperatures(T []int) []int {
	length := len(T)
	result := make([]int, length)
	stackWarmer := []int{}
	for i := length - 1; i >= 0; i-- {
		temp := T[i]
		// for j := length - 1; j > i; j-- {
		// 	if T[j] > temp {
		// 		stackWarmer = append(stackWarmer, j)
		// 	}
		// }
		// if len(stackWarmer) == 0 {
		// 	result[i] = 0
		// 	continue
		// }
		// result[i] = stackWarmer[len(stackWarmer)-1] - i
		for len(stackWarmer) > 0 && temp >= T[stackWarmer[len(stackWarmer)-1]] {
			log.Println("con co", stackWarmer)
			stackWarmer = stackWarmer[:len(stackWarmer)-1]
			log.Println("con co 2", stackWarmer)
		}
		if len(stackWarmer) > 0 {
			result[i] = stackWarmer[len(stackWarmer)-1] - i
		}
		stackWarmer = append(stackWarmer, i)
	}
	return result
}

func DailyTemperaturesSolution() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	log.Println(dailyTemperatures(T))
}

// EVALUATE REVERSE POLISH NOTATION

// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, /. Each operand may be an integer or another expression.

// Note:
// Division between two integers should truncate toward zero.
// The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.

// Example 1:
// Input: ["2", "1", "+", "3", "*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

// Example 2:
// Input: ["4", "13", "5", "/", "+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

// Example 3:
// Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
// Output: 22
// Explanation:
//   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

func evalRPN(tokens []string) int {
	operatorsMap := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	evalStack := []string{}
	for _, token := range tokens {
		if _, ok := operatorsMap[token]; ok {
			if len(evalStack) >= 2 {
				operand1, er1 := strconv.Atoi(evalStack[len(evalStack)-1])
				operand2, er2 := strconv.Atoi(evalStack[len(evalStack)-2])
				if er1 == nil && er2 == nil {
					res := 0
					switch token {
					case "+":
						res = (operand2 + operand1)
					case "-":
						res = (operand2 - operand1)
					case "*":
						res = (operand2 * operand1)
					case "/":
						res = (operand2 / operand1)
					}
					evalStack = evalStack[:len(evalStack)-2]
					resString := strconv.Itoa(res)
					evalStack = append(evalStack, resString)
					continue
				}
			}
		}
		evalStack = append(evalStack, token)
	}
	result, _ := strconv.Atoi(evalStack[0])
	return result
}

func EvalRPNSolution() {
	tokens := []string{"2", "1", "+", "3", "*"}
	log.Println(evalRPN(tokens))
}

// DFS - TEMPLATE I
// TEMPLATE - RECURSION
// /*
//  * Return true if there is a path from cur to target.
//  */
//  boolean DFS(Node cur, Node target, Set<Node> visited) {
//     return true if cur is target;
//     for (next : each neighbor of cur) {
//         if (next is not in visited) {
//             add next to visted;
//             return true if DFS(next, target, visited) == true;
//         }
//     }
//     return false;
// }

// NUMBER OF ISLANDS
// Given an m x n 2d grid map of '1's (land) and '0's (water), return the number of islands.

// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

// Example 1:
// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1

// Example 2:
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3

func numIslandDFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	visitedNode := make(map[string]bool)
	type nodeItem struct {
		row int
		col int
	}
	landNode := []nodeItem{}
	islands := 0
	for r, lines := range grid {
		for c, value := range lines {
			if string(value) == "1" {
				landNode = append(landNode, nodeItem{r, c})
			}
		}
	}
	for _, v := range landNode {
		if _, ok := visitedNode[fmt.Sprintf("%d-%d", v.row, v.col)]; ok {
			continue
		}
		count := 0
		numIslandRecursive(v.row, v.col, grid, visitedNode, &count)
		if count > 0 {
			islands++
		}
	}
	log.Println("con co", islands)
	return islands
}

func numIslandRecursive(r int, c int, grid [][]byte, visitedNode map[string]bool, count *int) {
	if r < 0 || c < 0 || r > len(grid)-1 || c > len(grid[0])-1 {
		return
	}
	if string(grid[r][c]) == "0" {
		return
	}
	*count++
	nextRow := r + 1
	prevRow := r - 1
	nextCol := c + 1
	prevCol := c - 1
	if _, ok := visitedNode[fmt.Sprintf("%d-%d", prevRow, c)]; !ok {
		visitedNode[fmt.Sprintf("%d-%d", prevRow, c)] = true
		numIslandRecursive(prevRow, c, grid, visitedNode, count)
	}
	if _, ok := visitedNode[fmt.Sprintf("%d-%d", nextRow, c)]; !ok {
		visitedNode[fmt.Sprintf("%d-%d", nextRow, c)] = true
		numIslandRecursive(nextRow, c, grid, visitedNode, count)
	}
	if _, ok := visitedNode[fmt.Sprintf("%d-%d", r, nextCol)]; !ok {
		visitedNode[fmt.Sprintf("%d-%d", r, nextCol)] = true
		numIslandRecursive(r, nextCol, grid, visitedNode, count)
	}
	if _, ok := visitedNode[fmt.Sprintf("%d-%d", r, prevCol)]; !ok {
		visitedNode[fmt.Sprintf("%d-%d", r, prevCol)] = true
		numIslandRecursive(r, prevCol, grid, visitedNode, count)
	}
}

// CLONE GRAPH

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node.Val]; ok {
		return v
	}
	clonedNode := &Node{Val: node.Val, Neighbors: []*Node{}}
	visited[clonedNode.Val] = clonedNode

	for _, n := range node.Neighbors {
		log.Println("con co", n)
		clone := clone(n, visited)
		if clone != nil {
			clonedNode.Neighbors = append(clonedNode.Neighbors, clone)
		}
	}
	return clonedNode
}

func ImplementCloneGraph() {
	node1 := Node{
		Val:       1,
		Neighbors: []*Node{},
	}
	node2 := Node{
		Val:       2,
		Neighbors: []*Node{},
	}
	node3 := Node{
		Val:       3,
		Neighbors: []*Node{},
	}
	node4 := Node{
		Val:       4,
		Neighbors: []*Node{},
	}
	node1.Neighbors = append(node1.Neighbors, &node2, &node4)
	node2.Neighbors = append(node2.Neighbors, &node1, &node3)
	node3.Neighbors = append(node3.Neighbors, &node2, &node4)
	node4.Neighbors = append(node4.Neighbors, &node1, &node3)

	cloned := cloneGraph(&node1)
	log.Println("con co be be", cloned)
}

// TARGET SUM
// You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

// // Find out how many ways to assign symbols to make sum of integers equal to target S.

// Constraints:
// The length of the given array is positive and will not exceed 20.
// The sum of elements in the given array will not exceed 1000.
// Your output answer is guaranteed to be fitted in a 32-bit integer.

// Example 1:
// Input: nums is [1, 1, 1, 1, 1], S is 3.
// Output: 5

// Explanation:
// -1+1+1+1+1 = 3
// +1-1+1+1+1 = 3
// +1+1-1+1+1 = 3
// +1+1+1-1+1 = 3
// +1+1+1+1-1 = 3
// There are 5 ways to assign symbols to make the sum of nums be target 3.

// example2
// nums = [29,6,7,36,30,28,35,48,20,44,40,2,31,25,6,41,33,4,35,38]
// S = 35

// example2
// nums = [42,24,30,14,38,27,12,29,43,42,5,18,0,1,12,44,45,50,21,47]
// S = 38

func findTargetSumWays(nums []int, S int) int {
	visited := map[string]int{}
	return sumWays(nums, 0, 0, S, visited)
}

func sumWays(nums []int, index int, sum int, S int, visited map[string]int) int {
	if index == len(nums) {
		if sum == S {
			return 1
		}
		return 0
	}
	if v, ok := visited[fmt.Sprintf("%d-%d", index, sum)]; ok {
		return v
	}
	add := sumWays(nums, index+1, sum+nums[index], S, visited)
	subtract := sumWays(nums, index+1, sum-nums[index], S, visited)
	visited[fmt.Sprintf("%d-%d", index, sum)] = add + subtract
	return visited[fmt.Sprintf("%d-%d", index, sum)]
}

func FindTargetSumWaysSolution() {
	nums := []int{42, 24, 30, 14, 38, 27, 12, 29, 43, 42, 5, 18, 0, 1, 12, 44, 45, 50, 21, 47}
	S := 38
	log.Println("con co", findTargetSumWays(nums, S))
}

// DFS - template II
/*
 * Return true if there is a path from cur to target.
 */
//  boolean DFS(int root, int target) {
//     Set<Node> visited;
//     Stack<Node> stack;
//     add root to stack;
//     while (s is not empty) {
//         Node cur = the top element in stack;
//         remove the cur from the stack;
//         return true if cur is target;
//         for (Node next : the neighbors of cur) {
//             if (next is not in visited) {
//                 add next to visited;
//                 add next to stack;
//             }
//         }
//     }
//     return false;
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []int{}
	node := root
	for node != nil {
		if node.Left == nil {
			stack = append(stack, node.Val)
			node = node.Right
		}
		node = node.Left
	}
	return stack
}

func ImplementDFSInorderTraversal() {
	node1 := TreeNode{1, nil, nil}
	node2 := TreeNode{2, nil, nil}
	node3 := TreeNode{3, nil, nil}
	node1.Right = &node2
	node2.Left = &node3
	log.Println("con co", inorderTraversal(&node1))
}
