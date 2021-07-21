package treesandgraphs

import "fmt"

func RunExampleOfGetPathWithSums() {
	tree := BinaryTree{}
	tree.Insert(NewBinaryTreeNode(1))
	tree.Insert(NewBinaryTreeNode(2))
	tree.Insert(NewBinaryTreeNode(3))
	tree.Insert(NewBinaryTreeNode(4))
	tree.Insert(NewBinaryTreeNode(5))
	tree.Insert(NewBinaryTreeNode(6))
	tree.Insert(NewBinaryTreeNode(7))
	tree.Insert(NewBinaryTreeNode(-1))
	tree.Insert(NewBinaryTreeNode(-3))
	tree.Insert(NewBinaryTreeNode(-2))

	fmt.Printf("\nNumber:%d", GetPathWithSums(tree.Root, []int{}, 3))
}


func GetPathWithSums(root *BinaryTreeNode, ancestors []int, lookup int) (numberOfOccurrence  int) {
	if root == nil{
		return
	}
	pathToCurrentNode := append(ancestors, root.Data)
	numberOfOccurrence += GetNumberOfTimesSumOfSequencesEqualToLookup(pathToCurrentNode, lookup)
	numberOfOccurrence += GetPathWithSums(root.LeftChild, pathToCurrentNode, lookup)
	numberOfOccurrence += GetPathWithSums(root.RightChild, pathToCurrentNode, lookup)
	return numberOfOccurrence

}

func GetNumberOfTimesSumOfSequencesEqualToLookup(path []int, lookup int) (count int) {
	sum := 0
	for i := len(path) - 1 ; i >= 0 ; i--{
		sum += path[i]
		if sum == lookup{
			count++
		}
	}
	fmt.Printf("\n%v\tcount:%d", path, count)
	return count
}