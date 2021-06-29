package main

import (
	"fmt"
	"github.com/eternity-wings/dev_talks/treesandgraphs"
	"github.com/eternity-wings/dev_talks/treesandgraphs/bstvalidator"
)

func main() {
	bst := treesandgraphs.BinarySearchTree{}
	bst.Add(10)
	bst.Add(5)
	bst.Add(11)
	bst.Add(3)
	bst.Add(6)
	bst.Add(1)
	//
	//treesandgraphs.PrintBreathNodes(bst)
	fmt.Printf("Is tree valid:%v", bstvalidator.IsTreeValid(bst))



}
