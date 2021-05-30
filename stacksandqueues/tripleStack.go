package stacksandqueues

import "fmt"

const NumberOfInternalStacks = 3

type TripleStack struct {
	Stack map[int]int
}

func CreateNewTripleStack() (stack TripleStack) {
	stack.Stack = make(map[int]int)
	return stack
}

func (t *TripleStack) DisplayStack(stackId int) {
	for index := t.GetStackTopIndex(stackId) - NumberOfInternalStacks; index > NumberOfInternalStacks; index -= NumberOfInternalStacks {
		fmt.Printf("%d ->", t.Stack[index])
	}
	fmt.Println()
}

func (t *TripleStack) Pop(stackId int) (result int) {
	stackSize := t.GetStackSize(stackId)
	if stackSize == 0 {
		return -1
	}
	topIndex := t.GetStackTopIndex(stackId)
	topItemIndex := topIndex - NumberOfInternalStacks
	result = t.Stack[topItemIndex]
	t.Stack[topItemIndex] = -1
	t.Stack[stackId]--
	return result
}

func (t *TripleStack) Push(stackId int, value int) {
	topIndex := t.GetStackTopIndex(stackId)
	t.Stack[topIndex] = value
	t.Stack[stackId]++
}

//
// [1][2][3] ([4]  [5]  [6])  ([7] [8] [9]) ([10][11][12]) ([13][14][15])
//     1      500  700         900
func (t *TripleStack) GetStackTopIndex(stackId int) int {
	stackSize := t.GetStackSize(stackId)
	return (stackSize * NumberOfInternalStacks) + (NumberOfInternalStacks + 1)
}

func (t *TripleStack) GetStackSize(stackId int) int {
	return t.Stack[stackId]
}
