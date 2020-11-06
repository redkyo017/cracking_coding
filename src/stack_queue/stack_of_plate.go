package stack_queue

import (
	"errors"
	"log"
)

type StackWithThreshold struct {
	Values    []int
	Threshold int
}

func (s *StackWithThreshold) initStack(threshold int) {
	s.Threshold = threshold
	s.Values = []int{}
}
func (s *StackWithThreshold) push(value int) error {
	if s.isFull() {
		return errors.New("this stack is full")
	}
	s.Values = append(s.Values, value)
	return nil
}
func (s *StackWithThreshold) isFull() bool {
	return s.Threshold == len(s.Values)
}
func (s *StackWithThreshold) isEmpty() bool {
	return len(s.Values) == 0
}
func (s *StackWithThreshold) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("this stack is empty")
	}
	v := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return v, nil
}

func (s *StackWithThreshold) peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("this stack is empty")
	}
	return s.Values[len(s.Values)-1], nil
}

type SetOfStacks struct {
	threshold int
	stacks    []*StackWithThreshold
}

func (ss *SetOfStacks) initSS(threshold int) {
	ss.threshold = threshold
	ss.stacks = []*StackWithThreshold{}
}
func (ss *SetOfStacks) getLastStack() (*StackWithThreshold, error) {
	if len(ss.stacks) == 0 {
		return nil, errors.New("all stack is empty")
	}
	return ss.stacks[len(ss.stacks)-1], nil
}

func (ss *SetOfStacks) push(value int) {
	stack, _ := ss.getLastStack()
	if stack == nil || stack.isFull() {
		stack = &StackWithThreshold{}
		stack.initStack(ss.threshold)
		ss.stacks = append(ss.stacks, stack)
	}
	stack.push(value)
}

func (ss *SetOfStacks) peek() (int, error) {
	if len(ss.stacks) == 0 {
		return 0, errors.New("stack is empty")
	}
	stack, _ := ss.getLastStack()
	if stack == nil || stack.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	return stack.peek()
}

func (ss *SetOfStacks) isEmpty() bool {
	stack, _ := ss.getLastStack()
	return stack == nil || stack.isEmpty()
}

func (ss *SetOfStacks) ShiftLeft(index int, removeTop bool) int {
	stack := ss.stacks[index]
	var remove_item int
	if removeTop {
		remove_item, _ = stack.pop()
	} else {
		remove_item = stack.Values[0]
		stack.Values = stack.Values[1:]
	}

	if stack.isEmpty() {
		ss.stacks = append(ss.stacks[0:index], ss.stacks[index+1:]...)
	} else if len(ss.stacks) > index+1 {
		v := ss.ShiftLeft(index+1, false)
		stack.push(v)
	}
	return remove_item
}

func (ss *SetOfStacks) PopAt(index int) int {
	return ss.ShiftLeft(index, true)
}

func (ss *SetOfStacks) pop() (int, error) {
	if len(ss.stacks) == 0 {
		return 0, errors.New("stack is empty")
	}
	stack, _ := ss.getLastStack()
	if stack == nil {
		return 0, errors.New("stack is not exist")
	}
	value, _ := stack.pop()
	if stack.isEmpty() {
		ss.stacks = ss.stacks[:len(ss.stacks)-1]
	}
	return value, nil
}

func ImplementSetOfStacks() {
	ss := SetOfStacks{}
	ss.initSS(3)
	ss.push(1)
	ss.push(2)
	ss.push(3)
	ss.push(4)
	ss.push(5)
	ss.push(6)
	ss.push(7)
	ss.push(8)
	p := ss.PopAt(2)
	log.Println("con meo", p)
	log.Println("con heo", len(ss.stacks))
	for _, s := range ss.stacks {
		log.Println("con co", s)
	}
	// p, _ := ss.pop()
	// log.Println("con co", p)
}
