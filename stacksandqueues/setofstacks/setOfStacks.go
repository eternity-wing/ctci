package setofstacks

import "fmt"

const Threshold = 3

type SetOfStacks struct {
	Top *StackNode
}

func (s *SetOfStacks) Push(value int) {
	if s.Top == nil || s.Top.Stack.size == Threshold {
		s.Top = &StackNode{
			NextStack: s.Top,
		}
	}
	s.Top.Stack.Push(value)
}

func (s *SetOfStacks) Pop() (value *int) {
	topStack := s.Top
	if topStack != nil && topStack.Stack.size == 0 {
		topStack = topStack.NextStack
	}

	if topStack == nil {
		return value
	}
	value = topStack.Stack.Pop()
	if topStack.Stack.size == 0 {
		s.Top = topStack.NextStack
	}

	return value
}

func (s *SetOfStacks) Display() {
	for stack := s.Top; stack != nil; stack = stack.NextStack {
		stack.Stack.Display()
	}
	fmt.Println()
}
