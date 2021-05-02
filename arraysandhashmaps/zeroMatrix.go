package arraysandhashmaps

import "fmt"

func ConvertToZeroMatrix(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])

	var zeroRows = make(map[int]bool)
	var zeroColumns = make(map[int]bool)

	for row := 0 ; row < m ; row++{
		for col := 0 ; col < n ; col++{
			 if matrix[row][col] == 0{
				zeroRows[row] = true
				zeroColumns[col] = true
			}
		}
	}

	for row := 0 ; row < m ; row++{
		for col := 0 ; col < n ; col++{
			if _, ok := zeroRows[row]; ok{
				matrix[row][col] = 0
			}else if _, ok := zeroColumns[col]; ok {
				matrix[row][col] = 0
			}
		}
	}

	for row :=0 ; row < m ; row++{
		for col :=0 ; col < n ; col++{
			fmt.Printf("%d\t", matrix[row][col])
		}
		fmt.Printf("\n")
	}
}