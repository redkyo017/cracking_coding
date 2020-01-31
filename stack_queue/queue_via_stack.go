package stack_queue

import (
	"log"
)

type QueueViaStack struct {
	OldStack *StackSimple
	NewStack *StackSimple
}

func (qs *QueueViaStack) InitQueue() {
	qs.OldStack = &StackSimple{}
	qs.NewStack = &StackSimple{}
	qs.OldStack.initStack()
	qs.NewStack.initStack()
}
func (qs *QueueViaStack) Size() int {
	return qs.OldStack.Size + qs.NewStack.Size
}
func (qs *QueueViaStack) Add(value int) {
	qs.NewStack.push(value)
}

func (qs *QueueViaStack) ShiftStack() {
	if qs.OldStack.isEmpty() {
		for !qs.NewStack.isEmpty() {
			v, err := qs.NewStack.pop()
			if err == nil {
				qs.OldStack.push(v)
			}
		}
	}
}

func (qs *QueueViaStack) Peek() (int, error) {
	qs.ShiftStack()
	return qs.OldStack.peek()
}
func (qs *QueueViaStack) Pop() (int, error) {
	qs.ShiftStack()
	return qs.OldStack.pop()
}

func ImplementQueueViaStacks() {
	qs := &QueueViaStack{}
	qs.InitQueue()
	qs.Add(1)
	qs.Add(2)
	qs.Add(3)
	qs.Add(4)
	qs.Add(5)
	v, _ := qs.Pop()
	log.Println("pop queue", v)
	v, _ = qs.Peek()
	log.Println("peek queue", v)
	qs.Add(6)
	qs.Add(7)
	qs.Pop()
	v, _ = qs.Peek()
	log.Println("2 peek queue", v)
}
