package treesandgraphs


type BinarySearchTree struct {
	Root *BinaryTreeNode
}

func (b *BinarySearchTree) Add(value int)  {
	newNode := &BinaryTreeNode{
		Data: value,
	}

	if b.Root == nil{
		b.Root = newNode
		return
	}

	for node := b.Root; node != nil;{
		if node.Data == newNode.Data{
			return
		}
		if node.Data < newNode.Data && node.RightChild == nil {
			node.RightChild = newNode
			newNode.Parent = node
			return
		}else if node.Data < newNode.Data{
			node = node.RightChild
			continue
		}

		if node.Data > newNode.Data && node.LeftChild == nil {
			node.LeftChild = newNode
			newNode.Parent = node
			return
		}else if node.Data > newNode.Data{
			node = node.LeftChild
			continue
		}

	}

}

func (b *BinarySearchTree) IsBalance() bool  {
	if b.Root == nil{
		return true
	}
	balanceFactor := GetBSTNodeBalanceFactor(*b.Root)
	return balanceFactor <= 1 && balanceFactor >= -1
}

func (b *BinarySearchTree) PrintInOrder()  {
	PrintInOrder(b.Root)
}


