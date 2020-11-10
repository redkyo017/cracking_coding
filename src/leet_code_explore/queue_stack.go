package leet_code_explore

import (
	"fmt"
	"log"
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
	landNode := make(map[string]bool)
	// visitedNode := make(map[string]bool)
	type nodeItem struct {
		row int
		col int
	}
	nodeQueue := []nodeItem{}
	nodeQueue = append(nodeQueue, nodeItem{0, 0})
	step := 0
	for r, lines := range grid {
		for c, value := range lines {
			if string(value) == "1" {
				landNode[fmt.Sprintf("%d-%d", r, c)] = true
			}
		}
	}
	for len(nodeQueue) > 0 {
		step++
		size := len(nodeQueue)
		log.Println("con meo", size)
		for i := 0; i < size; i++ {
			log.Println("con heo", i)
			node := nodeQueue[0]
			row := node.row
			col := node.col
			log.Println("con co 0", nodeQueue)
			if row < len(grid) && col < len(grid[0])-1 && string(grid[row][col+1]) == "1" {
				nodeQueue = append(nodeQueue, nodeItem{row, col + 1})
				log.Println("con co 1", nodeQueue, string(grid[row][col+1]))
			}
			if col < len(grid) && row < len(grid)-1 && string(grid[row+1][col]) == "1" {
				nodeQueue = append(nodeQueue, nodeItem{row + 1, col})
				log.Println("con co 2", nodeQueue, string(grid[row+1][col]))
			}
			nodeQueue = nodeQueue[1:]
			log.Println("con co 3", nodeQueue)
			log.Println("-----")
		}
	}
	log.Println("con co", grid, step)
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
	numIsland(grid)
}
