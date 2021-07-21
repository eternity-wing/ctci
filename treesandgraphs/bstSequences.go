package treesandgraphs

import "fmt"

// TLR
// TRL

func RunExampleOfPrintBstSequencesOfTheTree() {
	tree := BinarySearchTree{}
	tree.Add(20)
	tree.Add(10)
	tree.Add(22)
	tree.Add(5)
	tree.Add(27)
	tree.Add(15)
	tree.Add(25)
	tree.Add(12)
	PrintBstSequencesOfTheTree(tree)
}

func PrintBstSequencesOfTheTree(tree BinarySearchTree) {
	fmt.Printf("\nResult:%+v", GetBstSequencesOfTheNode(tree.Root))
}

func GetBstSequencesOfTheNode(root *BinaryTreeNode) (sequence [][]int) {
	if root == nil {
		return sequence
	}

	if root.IsLeaf() {
		return [][]int{{root.Data}}
	}

	leftSequences := GetBstSequencesOfTheNode(root.LeftChild)
	rightSequences := GetBstSequencesOfTheNode(root.RightChild)
	sequence = append(sequence, CombineSequences(PrependItemToAllTermsOfTheSequence(root.Data, leftSequences), rightSequences)...)
	sequence = append(sequence, CombineSequences(PrependItemToAllTermsOfTheSequence(root.Data, rightSequences), leftSequences)...)

	return sequence
}

func CombineSequences(leftSequences [][]int, rightSequences [][]int) (resultSequences [][]int) {
	for _, leftSequence := range leftSequences {
		if len(rightSequences) == 0 {
			resultSequences = append(resultSequences, leftSequence)
			continue
		}

		for _, rightSequence := range rightSequences {
			resultSequences = append(resultSequences, append(leftSequence, rightSequence...))
		}
	}
	return resultSequences
}

func PrependItemToAllTermsOfTheSequence(item int, sequences [][]int) (resultSequences [][]int) {
	for _, term := range sequences {
		resultSequences = append(resultSequences, append([]int{item}, term...))
	}
	return resultSequences
}
