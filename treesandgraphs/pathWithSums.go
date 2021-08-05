package treesandgraphs

import "fmt"

func RunExampleOfGetPathWithSums() {
	tree := BinaryTree{}
	six := &BinaryTreeNode{Data: 6}
	four := &BinaryTreeNode{Data: 4, RightChild: six}
	two := &BinaryTreeNode{Data: 2, RightChild: four}
	three := &BinaryTreeNode{Data: 3, RightChild: two}
	one := &BinaryTreeNode{Data: 1, RightChild: three}
	tree.Root = one

	fmt.Printf("\nNumber:%d", GetPathWithSums(tree.Root, []int{}, 6))
}

func GetPathWithSums(root *BinaryTreeNode, ancestors []int, lookup int) (numberOfOccurrence int) {
	if root == nil {
		return
	}
	pathToNode := append(ancestors, root.Data)                                            // o(h)
	numberOfOccurrence += GetNumberOfTimesSumOfSequencesEqualToLookup(pathToNode, lookup) //o(n)
	numberOfOccurrence += GetPathWithSums(root.LeftChild, pathToNode, lookup)
	numberOfOccurrence += GetPathWithSums(root.RightChild, pathToNode, lookup)
	return numberOfOccurrence

}

func GetNumberOfTimesSumOfSequencesEqualToLookup(path []int, lookup int) (count int) {
	sum := 0
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == lookup {
			count++
		}
	}
	return count
}
