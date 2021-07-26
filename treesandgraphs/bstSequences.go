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

	for index, sequence := range GetBstSequencesOfTheNode(tree.Root) {
		fmt.Printf("%d: %+v\n", index, sequence)
	}

}

func GetBstSequencesOfTheNode(root *BinaryTreeNode) (sequence [][]int) {
	if root == nil {
		return sequence
	}

	prefix :=[]int{root.Data}
	sequence = CombineSequences(prefix, GetBstSequencesOfTheNode(root.LeftChild), GetBstSequencesOfTheNode(root.RightChild))
	return sequence
}

func CombineSequences(prefix []int, leftSequences [][]int, rightSequences [][]int) (resultSequences [][]int) {
	if len(leftSequences) == 0 {
		leftSequences = [][]int{{}}
	}

	if len(rightSequences) == 0 {
		rightSequences = [][]int{{}}
	}


	for _, leftSequence := range leftSequences {

		for _, rightSequence := range rightSequences {
			resultSequences = append(resultSequences, Weave(prefix, leftSequence, rightSequence)...)
		}
	}
	return resultSequences
}


func Weave(prefix []int, leftSubtree []int, rightSubtree []int) (resultSequences [][]int) {
	if len(leftSubtree) == 0 || len(rightSubtree) == 0 {
		prefix = append(prefix, append(leftSubtree, rightSubtree...)...)
		return append(resultSequences, prefix)
	}

	headOfLeft, remainLeft := leftSubtree[0], leftSubtree[1:]
	resultSequences = append(resultSequences, Weave(append(prefix, headOfLeft), remainLeft, rightSubtree)...)

	headOfRight, remainRight := rightSubtree[0], rightSubtree[1:]
	resultSequences = append(resultSequences, Weave(append(prefix, headOfRight), leftSubtree, remainRight)...)

	return resultSequences
}

