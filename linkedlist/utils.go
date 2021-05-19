package linkedlist

import "fmt"

func AddNumericalLinkedLists(l1 BidirectionalLinkedList, l2 BidirectionalLinkedList) {
	var resultList = &BidirectionalLinkedList{}
	overflow := AddNodes(l1.Head, l2.Head, resultList)
	if overflow > 0 {
		resultNode := &Node{
			Next: resultList.Head,
			Data: overflow,
		}
		resultList.Head = resultNode
	}
	resultList.Display()
}

func AddNodes(firstNode *Node, secondNode *Node, resultList *BidirectionalLinkedList) (overflow int) {
	if firstNode == nil && secondNode == nil {
		return 0
	}

	firstDigit := 0
	secondDigit := 0
	var nextOfFirstNode *Node
	var nextOfSecondNode *Node
	if firstNode != nil {
		firstDigit = (firstNode.Data).(int)
		nextOfFirstNode = firstNode.Next
	}
	if secondNode != nil {
		secondDigit = (secondNode.Data).(int)
		nextOfSecondNode = secondNode.Next
	}
	overflow = AddNodes(nextOfFirstNode, nextOfSecondNode, resultList)
	sum := firstDigit + secondDigit + overflow
	resultDigit := sum % 10
	overflow = sum / 10

	resultNode := &Node{
		Next: resultList.Head,
		Data: resultDigit,
	}
	resultList.Head = resultNode
	fmt.Printf("\n%d + %d = %d\noverflow:%d result digit:%d\n", firstDigit, secondDigit, sum, overflow, resultDigit)

	return overflow

}
