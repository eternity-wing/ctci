package avltrees

type AVLNode struct {
	Parent     *AVLNode
	LeftChild  *AVLNode
	RightChild *AVLNode
	Height     int
	Data       int
}

func (a *AVLNode) GetBalanceFactor() int {
	return GetAvlNodeHeight(a.LeftChild) - GetAvlNodeHeight(a.RightChild)
}