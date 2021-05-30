package setofstacks

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type StackNode struct {
	NextStack *StackNode
	Stack     Stack
}

type Stack struct {
	Top  *Node
	size int
}

func (s *Stack) Push(value int) {
	Node := Node{
		Data: value,
		Next: s.Top,
	}
	s.Top = &Node
	s.size++
}

func (s *Stack) Pop() (value *int) {
	if s.Top != nil {
		value = &s.Top.Data
		s.Top = s.Top.Next
		s.size--
	}
	return value
}
func (s *Stack) Display() {
	for node := s.Top; node != nil; node = node.Next {
		fmt.Printf("%v ->", node.Data)
	}
	fmt.Println()
}
