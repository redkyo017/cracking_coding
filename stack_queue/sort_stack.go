package stack_queue

import "log"

func SortStack(s *StackSimple) {
	tempStack := &StackSimple{}
	tempStack.initStack()

	for !s.isEmpty() {
		tmp, _ := s.pop()
		p, _ := tempStack.peek()
		for !tempStack.isEmpty() && p > tmp {
			v, _ := tempStack.pop()
			s.push(v)
		}
		tempStack.push(tmp)
	}

	for !tempStack.isEmpty() {
		v, _ := tempStack.pop()
		s.push(v)
	}
}

func ImplementSortStack() {
	stack := &StackSimple{}
	stack.initStack()
	stack.push(5)
	stack.push(7)
	stack.push(12)
	stack.push(8)
	stack.push(1)
	stack.push(3)
	log.Println("con heo be be", stack)
	SortStack(stack)
	log.Println("con co be be", stack)
}
