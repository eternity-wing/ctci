package treesandgraphs

type BinaryTreeNode struct {
	Parent                *BinaryTreeNode
	LeftChild             *BinaryTreeNode
	RightChild            *BinaryTreeNode
	Data                  int
	NumberOfGrandChildren int
}

func (b *BinaryTreeNode) IsLeaf() bool {
	return b.RightChild == nil && b.LeftChild == nil
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{Data: data}
}


func (b *BinaryTreeNode) IncrementAncestorsNumberOfGrandChildren()  {
	if b == nil{
		return
	}
	b.NumberOfGrandChildren++
	b.Parent.IncrementAncestorsNumberOfGrandChildren()
}

func (b *BinaryTreeNode) DecrementAncestorsNumberOfGrandChildren()  {
	if b == nil{
		return
	}
	b.NumberOfGrandChildren--
	b.Parent.IncrementAncestorsNumberOfGrandChildren()
}