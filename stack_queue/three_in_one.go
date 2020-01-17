package stack_queue

import "log"

type FixedMultiStack struct {
	NumberOfStacks int
	StackCapacity  int
	Sizes          []int
	Values         []int
}

func (fms *FixedMultiStack) Show() {
	log.Println("con co be be", fms)
}
