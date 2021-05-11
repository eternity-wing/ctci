package main

import (
	"github.com/eternity-wings/dev_talks/linkedlist"
)

func main() {
	firstArray := []int{7, 1, 6}
	firstList := linkedlist.BidirectionalLinkedList{}
	for _, value := range firstArray {
		firstList.InsertItem(value)
	}

	secondArray := []int{5, 9, 2, 3, 3, 1}
	secondList := linkedlist.BidirectionalLinkedList{}
	for _, value := range secondArray {
		secondList.InsertItem(value)
	}

	_ = firstList.Add(secondList)

}
