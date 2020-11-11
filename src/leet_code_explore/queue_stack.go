package leet_code_explore

import (
	"fmt"
	"log"
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
	// grid := [][]byte{
	// 	[]byte("11000"),
	// 	[]byte("11000"),
	// 	[]byte("00100"),
	// 	[]byte("00011"),
	// }
	// grid := [][]byte{
	// 	[]byte("11110"),
	// 	[]byte("11010"),
	// 	[]byte("11000"),
	// 	[]byte("00000"),
	// }
	grid := [][]byte{
		[]byte("111"),
		[]byte("010"),
		[]byte("111"),
	}
	// grid := [][]byte{
	// 	[]byte("10111"),
	// 	[]byte("10101"),
	// 	[]byte("11101"),
	// }
	numIsland(grid)
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
