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

// IMPLEMENT FLEXIBLE SIZE STACK
type StackInfo struct {
	Start    int
	Size     int
	Capacity int
}

func (st *StackInfo) IsFull() bool {
	return st.Size == st.Capacity
}

func (st *StackInfo) IsEmpty() bool {
	return st.Size == 0
}

func (st *StackInfo) InitStack(start, capacity int) {
	st.Start = start
	st.Capacity = capacity
}

type MultiStack struct {
	StackInfo []*StackInfo
	Values    []int
}

func (ms *MultiStack) InitMultiStack(numberOfStacks, defaultSize int) {
	for i := 0; i < numberOfStacks; i++ {
		st := &StackInfo{}
		st.Start = i * defaultSize
		st.Capacity = defaultSize
		ms.StackInfo = append(ms.StackInfo, st)
	}
	ms.Values = make([]int, numberOfStacks*defaultSize)
}

func (ms *MultiStack) IsWithinStackCapacity(index int, st *StackInfo) bool {
	if index < 0 || index >= len(ms.Values) {
		return false
	}
	var contiguousIndex int
	if index < st.Start {
		contiguousIndex = st.Start + len(ms.Values)
	} else {
		contiguousIndex = index
	}
	end := st.Start + st.Capacity
	return st.Start <= contiguousIndex && contiguousIndex <= end
}

func (ms *MultiStack) LastCapacityIndex(st *StackInfo) int {
	return ms.AdjustIndex(st.Start + st.Capacity - 1)
}

func (ms *MultiStack) LastElementIndex(st *StackInfo) int {
	return ms.AdjustIndex(st.Start + st.Size - 1)
}

func (ms *MultiStack) NumberOfElements() int {
	var sizes int
	for _, st := range ms.StackInfo {
		sizes += st.Size
	}
	return sizes
}

func (ms *MultiStack) AllStackFull() bool {
	return ms.NumberOfElements() == len(ms.Values)
}

func (ms *MultiStack) AdjustIndex(index int) int {
	max := len(ms.Values)
	return ((index % max) + max) % max
}

func (ms *MultiStack) NextIndex(index int) int {
	return ms.AdjustIndex(index + 1)
}

func (ms *MultiStack) PreviousIndex(index int) int {
	return ms.AdjustIndex(index - 1)
}

func (ms *MultiStack) Peek(stackNum int) int {
	st := ms.StackInfo[stackNum]
	return ms.Values[ms.LastElementIndex(st)]
}

func (ms *MultiStack) Pop(stackNum int) (int, error) {
	st := ms.StackInfo[stackNum]
	if st.IsEmpty() {
		return 0, errors.New("this stack is empty")
	}
	value := ms.Values[ms.LastElementIndex(st)]
	ms.Values[ms.LastElementIndex(st)] = 0
	st.Size--
	return value, nil
}

func (ms *MultiStack) Shift(stackNum int) {
	log.Println("shifting stack")
	st := ms.StackInfo[stackNum]
	if st.Size >= st.Capacity {
		nextStack := (stackNum + 1) % len(ms.StackInfo)
		ms.Shift(nextStack)
		st.Capacity++
	}
	idx := ms.LastCapacityIndex(st)
	for ms.IsWithinStackCapacity(idx, st) {
		ms.Values[idx] = ms.Values[ms.PreviousIndex(idx)]
		idx = ms.PreviousIndex(idx)
	}

	ms.Values[st.Start] = 0
	st.Start = ms.NextIndex(st.Start)
	st.Capacity--
}

func (ms *MultiStack) Expand(stackNum int) {
	ms.Shift((stackNum + 1) % len(ms.StackInfo))
	ms.StackInfo[stackNum].Capacity++
}

func (ms *MultiStack) Push(stackNum int, value int) error {
	if ms.AllStackFull() {
		return errors.New("all stack are full")
	}

	st := ms.StackInfo[stackNum]
	if st.IsFull() {
		ms.Expand(stackNum)
	}
	st.Size++
	ms.Values[ms.LastElementIndex(st)] = value

	return nil
}

func ImplementStackByOneArray() {
	// fms := &FixedMultiStack{}
	// fms.Init3Stack(5)
	// log.Println("con co", fms)
	// fms.Push(1, 9)
	// fms.Push(1, 8)
	// log.Println("con heo", fms)
	// fms.Pop(1)
	// log.Println("con meo", fms)

	ms := &MultiStack{}
	ms.InitMultiStack(3, 5)
	log.Println("con co", ms)
	ms.Push(2, 1)
	ms.Push(2, 2)
	ms.Push(2, 3)
	ms.Push(2, 4)
	ms.Push(2, 5)
	ms.Push(2, 6)
	log.Println("con heo", ms)
}
