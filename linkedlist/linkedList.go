package linkedlist

import "fmt"

type BidirectionalLinkedList struct {
	Head   *Node
	Tail   *Node
	length int
}

func (b *BidirectionalLinkedList) Len() int {
	return b.length
}

func (b *BidirectionalLinkedList) Display() {
	for node := b.Head; node != nil; node = node.Next {
		fmt.Printf("%v ->", node.Data)
	}
	fmt.Println()
}

func (b *BidirectionalLinkedList) DeleteNode(node *Node) {
	if node == b.Head {
		b.Head = node.Next
	}
	if node == b.Tail {
		b.Tail = node.Previous
	}

	prev := node.Previous
	next := node.Next
	if next != nil {
		next.Previous = prev
	}

	if prev != nil {
		prev.Next = next
	}
}

func (b *BidirectionalLinkedList) SinglyDeleteNode(node *Node) {
	if node == b.Head {
		b.Head = node.Next
	}

	next := node.Next
	if next != nil {
		node.Data = next.Data
		node.Next = next.Next
	} else {
		node = nil
	}

}

func (b *BidirectionalLinkedList) InsertItem(data interface{}) {
	node := &Node{
		Data:     data,
		Previous: b.Tail,
	}
	if b.Head == nil {
		b.Head = node
		b.Tail = node
	} else {
		b.Tail.Next = node
		b.Tail = node
	}
	b.length++
}

func (b *BidirectionalLinkedList) InsertNode(node *Node) {
	if b.Head == nil {
		b.Head = node
		b.Tail = node
	} else {
		b.Tail.Next = node
		b.Tail = node
	}
	b.length++
}

func (b *BidirectionalLinkedList) DeleteDuplicates() {
	for node := b.Head; node != nil; node = node.Next {
		for runner := node.Next; runner != nil; runner = runner.Next {
			if runner.Data == node.Data {
				b.DeleteNode(runner)
			}
		}
	}
}

func (b *BidirectionalLinkedList) PrintKth(k int) {
	kthNode := b.Head
	distance := 1
	for runner := kthNode.Next; runner != nil; runner = runner.Next {
		distance++
		if distance > k {
			distance--
			kthNode = kthNode.Next
		}
	}

	fmt.Printf("%v is the %dth last elementh", kthNode.Data, k)
}

func (b *BidirectionalLinkedList) Partition(pivot int) {
	pivotNode := b.Head
	for runner := pivotNode; runner != nil; runner = runner.Next {
		value := (runner.Data).(int)
		if value < pivot {
			runner.Data = pivotNode.Data
			pivotNode.Data = value
			pivotNode = pivotNode.Next
		}
	}
}

func (b *BidirectionalLinkedList) Add(secondList BidirectionalLinkedList) (result BidirectionalLinkedList) {
	firstPointer := b.Head
	secondPointer := secondList.Head
	overflow := 0
	for firstPointer != nil || secondPointer != nil {
		firstDigit := 0
		secondDigit := 0
		if firstPointer != nil {
			firstDigit = (firstPointer.Data).(int)
			firstPointer = firstPointer.Next
		}
		if secondPointer != nil {
			secondDigit = (secondPointer.Data).(int)
			secondPointer = secondPointer.Next
		}
		sum := firstDigit + secondDigit + overflow
		resultDigit := sum % 10
		overflow = sum / 10
		resultNode := &Node{
			Next: result.Head,
			Data: resultDigit,
		}
		result.Head = resultNode
	}
	result.Display()
	return result

}

func (b *BidirectionalLinkedList) PrintIntersections(secondList BidirectionalLinkedList) {
	for l1Node := b.Head; l1Node != nil; l1Node = l1Node.Next {
		for l2Node := secondList.Head; l2Node != nil; l2Node = l2Node.Next {
			if l2Node == l1Node {
				fmt.Printf("%v ->", l2Node.Data)
			}
		}
	}
}

func (b *BidirectionalLinkedList) IsPalindrome() bool {
	reverseList := BidirectionalLinkedList{}
	for node := b.Head; node != nil; node = node.Next {
		reverseHead := reverseList.Head
		newNode := &Node{
			Data: node.Data,
			Next: reverseHead,
		}
		reverseList.Head = newNode
	}

	l1Node := b.Head
	reverseNode := reverseList.Head
	for l1Node != nil {
		if l1Node.Data != reverseNode.Data {
			return false
		}
		l1Node = l1Node.Next
		reverseNode = reverseNode.Next
	}

	return true

}

func (b *BidirectionalLinkedList) LoopBeginningDetector() {
	loopBeginning := b.GetBeginningDetector()

	if loopBeginning == nil {
		fmt.Printf("\nThere is no loop")
	} else {
		fmt.Printf("\nBeginning node:%+v", loopBeginning)
	}
}

func (b *BidirectionalLinkedList) GetBeginningDetector() (LoopBeginning *Node) {
	loopCollisionSpot := b.GetLoopCollisionSpot()

	if loopCollisionSpot == nil {
		return nil
	}
	for turtle := b.Head; ; {
		isLoopBeginningSpot := turtle == loopCollisionSpot
		if isLoopBeginningSpot {
			LoopBeginning = turtle
			return LoopBeginning
		}
		turtle = turtle.Next
		loopCollisionSpot = loopCollisionSpot.Next
	}

}

func (b *BidirectionalLinkedList) GetLoopCollisionSpot() (CollisionSpot *Node) {
	turtle := b.Head
	rabbit := b.Head

	for !(turtle.Next == nil || rabbit.Next == nil || rabbit.Next.Next == nil) {
		turtle = turtle.Next
		rabbit = rabbit.Next.Next
		isCollisionSpot := turtle == rabbit
		if isCollisionSpot {
			CollisionSpot = rabbit
			return CollisionSpot
		}
	}
	return nil

}
