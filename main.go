package main

import (
	"github.com/eternity-wings/dev_talks/stacksandqueues/setofstacks"
)

func main() {

	firstArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	stacks := setofstacks.SetOfStacks{}
	for _, value := range firstArray {
		stacks.Push(value)
	}
	stacks.Pop()
	stacks.Pop()
	stacks.Display()
}
