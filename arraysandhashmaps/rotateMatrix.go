package arraysandhashmaps

import "fmt"

func RotateMatrix(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])
	var rotatedMatrix = make(map[int][]int, n)
	for i, row := range matrix{
		for j, element := range row{
			rotatedI, rotatedJ := GetRotatedIndex(i, j, m, n)
			if _, ok := rotatedMatrix[rotatedI]; !ok{
				rotatedMatrix[rotatedI] = make([]int, m)
			}
			rotatedMatrix[rotatedI][rotatedJ] = element
		}
	}

	for row :=0 ; row < n ; row++{
		for col :=0 ; col < m ; col++{
			fmt.Printf("%d\t", rotatedMatrix[row][col])
		}
		fmt.Printf("\n")
	}

}

func GetRotatedIndex(i int, j int, m int, n int) (int, int)  {
	rotatedI := n - (j + 1)
	rotatedJ := i
	return rotatedI, rotatedJ
}