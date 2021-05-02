package main

import (
	"github.com/eternity-wings/dev_talks/arraysandhashmaps"
)

func main()  {
	myMtr := [][]int{
		{0 , 2 ,  3},
		{4 , 0 ,  6},
		{7 , 8 ,  9},
		{10, 11, 12},
	}

	arraysandhashmaps.ConvertToZeroMatrix(myMtr)
}