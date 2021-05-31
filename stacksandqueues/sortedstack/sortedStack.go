package sortedstack

import "github.com/eternity-wings/dev_talks/stacksandqueues"

type SortedStack struct {
	stack stacksandqueues.Stack
}

func (s *SortedStack) Push(value int) {
	Node := stacksandqueues.Node{
		Data: value,
	}
	tempStack := stacksandqueues.Stack{}

	for nodeValue := s.Peek(); nodeValue != nil && *nodeValue < Node.Data; nodeValue = s.Peek() {
		tempStack.Push(*nodeValue)
		_ = s.stack.Pop()
	}
	s.stack.Push(value)

	for nodeValue := tempStack.Pop(); nodeValue != nil; nodeValue = tempStack.Pop() {
		s.stack.Push(*nodeValue)
	}
}

func (s *SortedStack) Pop() (value *int) {
	return s.stack.Pop()
}

func (s *SortedStack) Peek() (value *int) {
	if s.stack.Head != nil {
		value = &s.stack.Head.Data
	}
	return value
}

func (s *SortedStack) IsEmpty() bool {
	return s.stack.Head == nil
}

func (s *SortedStack) Display() {
	s.stack.Display()
}
