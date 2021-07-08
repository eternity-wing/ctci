package treesandgraphs

type BinaryTreeNode struct {
	Parent     *BinaryTreeNode
	LeftChild  *BinaryTreeNode
	RightChild *BinaryTreeNode
	Data       int
}


func (b *BinaryTreeNode) IsLeaf() bool  {
	return b.RightChild == nil && b.LeftChild == nil
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{Data: data}
}