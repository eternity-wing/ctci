package stacksandqueues

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type Stack struct {
	Head *Node
}

func (s *Stack) Push(value int) {
	Node := Node{
		Data: value,
		Next: s.Head,
	}
	s.Head = &Node
}

func (s *Stack) Pop() (value *int) {
	if s.Head != nil {
		value = &s.Head.Data
		s.Head = s.Head.Next
	}
	return value
}

func (s *Stack) Display() {
	for node := s.Head; node != nil; node = node.Next {
		fmt.Printf("%v ->", node.Data)
	}
	fmt.Println()
}
