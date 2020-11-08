package leet_code_explore

import "log"

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
