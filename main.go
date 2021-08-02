package main

import (
	"fmt"
	"github.com/eternity-wings/dev_talks/treesandgraphs"
	"github.com/eternity-wings/dev_talks/treesandgraphs/bstvalidator"
)

func main() {


	//treesandgraphs.RunExampleOfPrintBstSequencesOfTheTree()
	tree := treesandgraphs.BinarySearchTree{}
	six := &treesandgraphs.BinaryTreeNode{
		Data: 2,
		LeftChild:  &treesandgraphs.BinaryTreeNode{
			Data: 2,
		},
		RightChild:  &treesandgraphs.BinaryTreeNode{
			Data: 2,
		},
	}
	tree.Root = six
	fmt.Printf("\nIs valid:%v", bstvalidator.IsTreeValid(tree))

	//treesandgraphs.RunExampleOfGetPathWithSums()
	//treesandgraphs.PrintInOrder(tree.Root)
}
