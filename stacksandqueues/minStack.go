package stacksandqueues

import "fmt"

type MinStackNode struct {
	Next *MinStackNode
	Min  int
	Data int
}

type MinStack struct {
	Head *MinStackNode
}

func (m *MinStack) Push(value int) {
	Node := MinStackNode{
		Data: value,
		Min:  value,
		Next: m.Head,
	}
	if m.Head != nil && m.Head.Min < value {
		Node.Min = m.Head.Min
	}
	m.Head = &Node
}

func (m *MinStack) Pop() (value *int) {
	if m.Head != nil {
		value = &m.Head.Data
		m.Head = m.Head.Next
	}
	return value
}

func (m *MinStack) Min() (value *int) {
	if m.Head != nil {
		value = &m.Head.Min
	}
	return value
}

func (m *MinStack) Display() {
	for node := m.Head; node != nil; node = node.Next {
		fmt.Printf("%v ->", node.Data)
	}
	fmt.Println()
}
