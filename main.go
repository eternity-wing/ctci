package main

import (
	"fmt"
	"github.com/eternity-wings/dev_talks/treesandgraphs"
	"math/rand"
	"time"
)

func main() {
	Init()

	bst := treesandgraphs.BinarySearchTree{}
	root := &treesandgraphs.BinaryTreeNode{Data: 1}

	leftTwo := &treesandgraphs.BinaryTreeNode{
		Data: 2,
		RightChild: &treesandgraphs.BinaryTreeNode{
			Data: 4,
		},
		LeftChild: &treesandgraphs.BinaryTreeNode{
			Data: 3,
			LeftChild: &treesandgraphs.BinaryTreeNode{
				Data: 5,
				LeftChild: &treesandgraphs.BinaryTreeNode{
					Data: 6,
				},
			},
		},
	}


	rightTwo := &treesandgraphs.BinaryTreeNode{
		Data: 2,
		RightChild: &treesandgraphs.BinaryTreeNode{
			Data: 3,
		},
		LeftChild: &treesandgraphs.BinaryTreeNode{
			Data: 4,
			LeftChild: &treesandgraphs.BinaryTreeNode{
				Data: 5,
				LeftChild: &treesandgraphs.BinaryTreeNode{
					Data: 6,
				},
			},
		},
	}
	root.LeftChild = leftTwo
	root.RightChild = rightTwo

	bst.Root = root


	fmt.Printf("\n%+v", treesandgraphs.IsBalance(root, new(int)))
	//treesandgraphs.RunExampleOfPrintBstSequencesOfTheTree()

	//treesandgraphs.RunExampleOfGetPathWithSums()
	//treesandgraphs.PrintInOrder(tree.Root)
}

func Init() {
	rand.Seed(time.Now().UnixNano())
}
