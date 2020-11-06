package stack_queue

import (
	"errors"
	"log"
)

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
	MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

type StackWithMin struct {
	Values []int
	min    []int
}

func (st *StackWithMin) Peek() int {
	return st.Values[len(st.Values)-1]
}

func (st *StackWithMin) isEmpty() bool {
	return len(st.Values) == 0
}

func (st *StackWithMin) Min() int {
	if len(st.min) == 0 {
		return MaxInt
	}
	return st.min[len(st.min)-1]
}

func (st *StackWithMin) Push(value int) error {
	st.Values = append(st.Values, value)
	if value < st.Min() {
		st.min = append(st.min, value)
	}
	return nil
}

func (st *StackWithMin) Pop() (int, error) {
	if st.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	val := st.Values[len(st.Values)-1]
	st.Values = st.Values[:len(st.Values)-1]

	if st.Min() >= val {
		st.min = st.min[:len(st.min)-1]
	}

	return val, nil
}

func ShowStackWithMin() {
	st := &StackWithMin{}
	st.Push(3)
	st.Push(4)
	st.Push(2)
	st.Push(7)
	st.Push(1)
	log.Println("con co", st)
	st.Pop()
	st.Pop()
	log.Println("con heo", st)
}
