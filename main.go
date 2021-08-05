package main

import (
	"github.com/eternity-wings/dev_talks/treesandgraphs"
	"math/rand"
	"time"
)

func main() {
	Init()
	//treesandgraphs.RunExampleOfPrintBstSequencesOfTheTree()
	treesandgraphs.RunExampleOfGetPathWithSums()

	//treesandgraphs.RunExampleOfGetPathWithSums()
	//treesandgraphs.PrintInOrder(tree.Root)
}

func Init()  {
	rand.Seed(time.Now().UnixNano())
}
