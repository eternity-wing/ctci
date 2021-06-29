package avltrees

import (
	"fmt"
	"math"
)

type AVLTree struct {
	Root *AVLNode
}

func (a *AVLTree) Insert(value int) {
	a.Root = a.insertNode(a.Root, value)
}

func (a *AVLTree) insertNode(root *AVLNode, value int) *AVLNode {
	if root == nil {
		return &AVLNode{Data: value}
	}
	if root.Data < value {
		root.RightChild = a.insertNode(root.RightChild, value)
	} else {
		root.LeftChild = a.insertNode(root.LeftChild, value)
	}

	a.CalculateNodeHeight(root)
	a.BalanceTheNode(root)
	return root
}

func (a *AVLTree) BalanceTheNode(node *AVLNode)  {
	if a.IsLeftHeavy(node) {
		if a.GetBalanceFactor(node.LeftChild) < 0{
			fmt.Printf("left rotate\n")
		}
		fmt.Printf("right rotate\n")
	}else if a.IsRightHeavy(node){
		if a.GetBalanceFactor(node.RightChild) > 0{
			fmt.Printf("right rotate\n")
		}
		fmt.Printf("left rotate\n")
	}
}


func (a *AVLTree) GetBalanceFactor(node *AVLNode) int {
	if node == nil{
		return 0
	}
	return a.GetAvlNodeHeight(node.LeftChild) - a.GetAvlNodeHeight(node.RightChild)
}

func(a *AVLTree) GetAvlNodeHeight(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func (a *AVLTree) CalculateNodeHeight(node *AVLNode) {
	if node == nil {
		return
	}
	node.Height = int(math.Max(float64(a.GetAvlNodeHeight(node.LeftChild)), float64(a.GetAvlNodeHeight(node.RightChild)))) + 1
}

func (a *AVLTree) IsLeftHeavy(node *AVLNode) bool {
	return a.GetBalanceFactor(node) > 1
}

func (a *AVLTree) IsRightHeavy(node *AVLNode) bool {
	return a.GetBalanceFactor(node) < -1
}


func (a *AVLTree) PrintInOrder(node *AVLNode) {
	if node == nil {
		return
	}
	a.PrintInOrder(node.LeftChild)
	fmt.Printf("%d -> ", node.Data)
	a.PrintInOrder(node.RightChild)
}