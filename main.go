package main

import (
	"github.com/eternity-wings/dev_talks/linkedlist"
)

func main() {

	//array := []int{7 , 4 , 6 , 1 , 8 , 9 , 3 , 2}
	//l1 := linkedlist.BidirectionalLinkedList{}
	//for _, value := range array {
	//	l1.InsertItem(value)
	//}
	//l1.Display()
	//l1.Partition(4)
	//fmt.Println()
	//l1.Display()

	//array := []int{5, 9, 2, 3, 3, 1}
	//l1 := linkedlist.BidirectionalLinkedList{}
	//for _, value := range array {
	//	l1.InsertItem(value)
	//}
	//l2 := linkedlist.BidirectionalLinkedList{}
	//l2.Head = l1.Head.Next.Next
	//_ = l1.PrintIntersections(l2)

	//str := "abbaa"
	//l1 := linkedlist.BidirectionalLinkedList{}
	//for _, value := range str {
	//	l1.InsertItem(value)
	//}
	//fmt.Printf("\nIs Palindrome:%v", l1.IsPalindrome())

	//firstArray := []int{7, 1, 6, 1}
	//firstList := linkedlist.BidirectionalLinkedList{}
	//for _, value := range firstArray {
	//	firstList.InsertItem(value)
	//}
	//secondArray := []int{9, 1, 2}
	//secondList := linkedlist.BidirectionalLinkedList{}
	//for _, value := range secondArray {
	//	secondList.InsertItem(value)
	//}
	//linkedlist.AddNumericalLinkedLists(firstList, secondList)

	firstArray := []string{"A", "B", "C", "D", "E", "F", "G", "H", "C"}
	firstList := linkedlist.BidirectionalLinkedList{}
	for _, value := range firstArray {
		firstList.InsertItem(value)
	}
	firstList.Tail.Next = firstList.Head.Next.Next

	firstList.LoopBeginningDetector()

}
