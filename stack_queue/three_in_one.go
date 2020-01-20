package stack_queue

import (
	"errors"
	"log"
)

type FixedMultiStack struct {
	NumberOfStacks int
	StackCapacity  int
	Sizes          []int
	Values         []int
}

func (fms *FixedMultiStack) Init3Stack(stackSize int) {
	fms.NumberOfStacks = 3
	fms.StackCapacity = stackSize
	fms.Sizes = make([]int, fms.NumberOfStacks)
	fms.Values = make([]int, fms.NumberOfStacks*stackSize)
}

func (fms *FixedMultiStack) IsFull(stackNum int) bool {
	return fms.Sizes[stackNum] == fms.StackCapacity
}

func (fms *FixedMultiStack) IsEmpty(stackNum int) bool {
	return fms.Sizes[stackNum] == 0
}

func (fms *FixedMultiStack) IndexOfTop(stackNum int) int {
	offset := stackNum * fms.StackCapacity
	return offset + fms.Sizes[stackNum]
}

func (fms *FixedMultiStack) Push(stackNum, value int) error {
	if fms.IsFull(stackNum) {
		return errors.New("this stack is full")
	}
	topIndexStack := fms.IndexOfTop(stackNum)
	fms.Values[topIndexStack+1] = value
	fms.Sizes[stackNum]++
	return nil
}

func (fms *FixedMultiStack) Peek(stackNum int) (int, error) {
	if fms.IsEmpty(stackNum) {
		return 0, errors.New("this stack is empty")
	}
	return fms.Values[fms.IndexOfTop(stackNum)], nil
}

func (fms *FixedMultiStack) Pop(stackNum int) (int, error) {
	if fms.IsEmpty(stackNum) {
		return 0, errors.New("this stack is empty")
	}
	index := fms.IndexOfTop(stackNum)
	value := fms.Values[index]
	fms.Values[index] = 0
	fms.Sizes[stackNum]--
	return value, nil
}

func ImplementStackByOneArray() {
	fms := &FixedMultiStack{}
	fms.Init3Stack(5)
	log.Println("con co", fms)
	fms.Push(1, 9)
	fms.Push(1, 8)
	log.Println("con heo", fms)
	fms.Pop(1)
	log.Println("con meo", fms)
}
