package setofstacks

import "fmt"

type EnhancedSetOfStacks struct {
	stacks              map[int]Stack
	noneFilledStacksIds Stack
	numberOfStacks      int
}

func (e *EnhancedSetOfStacks) Push(value int) {
	topStack, isTopStackExist := e.stacks[e.numberOfStacks]
	acceptorStackIndex := e.numberOfStacks

	noneFullStack := e.noneFilledStacksIds.Pop()
	if noneFullStack != nil {
		acceptorStackIndex = *noneFullStack
	} else if isTopStackExist && topStack.size == Threshold {
		e.numberOfStacks++
		acceptorStackIndex = e.numberOfStacks
	}

	if _, isAcceptorStackExist := e.stacks[acceptorStackIndex]; !isAcceptorStackExist {
		e.stacks[acceptorStackIndex] = Stack{}
	}
	acceptorStack := e.stacks[acceptorStackIndex]
	acceptorStack.Push(value)
}

func (e *EnhancedSetOfStacks) Pop() (value *int) {
	for stackIndex := e.numberOfStacks; stackIndex >= 0; stackIndex-- {
		stack, isStackExist := e.stacks[stackIndex]
		if !isStackExist && stack.size == 0 {
			continue
		}
		if stackIndex != e.numberOfStacks {
			e.noneFilledStacksIds.Push(stackIndex)
		}
		return stack.Pop()
	}

	return value
}

func (e *EnhancedSetOfStacks) PopAt(index int) (value *int) {
	stack, isStackExist := e.stacks[index]
	if !isStackExist {
		return value
	}
	if index != e.numberOfStacks && stack.size > 0 {
		e.noneFilledStacksIds.Push(index)
	}
	return stack.Pop()
}

func (e *EnhancedSetOfStacks) Display() {
	for stackIndex := 0; stackIndex < e.numberOfStacks; stackIndex++ {
		stack := e.stacks[stackIndex]
		stack.Display()
	}
	fmt.Println()
}
