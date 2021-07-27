package checksubtree

import (
	"fmt"
	"github.com/eternity-wings/dev_talks/treesandgraphs"
)

func IsIncluded(bigTree treesandgraphs.BinaryTree, subTree treesandgraphs.BinaryTree) bool {
	return IsSubstring(GetStrictPreOrderPresentation(bigTree.Root), GetStrictPreOrderPresentation(subTree.Root))
}

func IsSubstring(str1 string, str2 string) bool {
	str2length := len(str2)

	str2Pointer := 0
	for str1Pointer, _ := range str1 {
		if str1[str1Pointer] == str2[str2Pointer] {
			str2Pointer++
		} else {
			str2Pointer = 0
		}
		if str2Pointer == str2length {
			return true
		}
	}
	return false
}

func GetStrictPreOrderPresentation(node *treesandgraphs.BinaryTreeNode) (preOrderPresentation string) {
	if node == nil {
		return "*"
	}
	return fmt.Sprintf("%d%s%s", node.Data, GetStrictPreOrderPresentation(node.LeftChild), GetStrictPreOrderPresentation(node.RightChild))
}

func RunExampleOfIsSubtree() {
	tree := treesandgraphs.BinaryTree{}
	tree.Insert(treesandgraphs.NewBinaryTreeNode(10))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(15))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(13))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(16))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(17))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(5))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(3))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(4))
	tree.Insert(treesandgraphs.NewBinaryTreeNode(6))

	subtree := treesandgraphs.BinaryTree{}
	node16 := treesandgraphs.NewBinaryTreeNode(16)
	node16.LeftChild = treesandgraphs.NewBinaryTreeNode(4)
	node16.RightChild = treesandgraphs.NewBinaryTreeNode(6)
	subtree.Insert(node16)

	fmt.Printf("\nIs Subtree:%+v", IsIncluded(tree, subtree))
}
