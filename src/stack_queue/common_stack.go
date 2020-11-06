package stack_queue

import "errors"

type StackSimple struct {
	Values []int
	Size   int
}

func (s *StackSimple) initStack() {
	s.Values = []int{}
	s.Size = 0
}
func (s *StackSimple) isEmpty() bool {
	return s.Size == 0 || len(s.Values) == 0
}
func (s *StackSimple) push(value int) {
	s.Values = append(s.Values, value)
	s.Size++
}
func (s *StackSimple) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	value := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	s.Size--
	return value, nil
}
func (s *StackSimple) peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	value := s.Values[len(s.Values)-1]
	return value, nil
}
