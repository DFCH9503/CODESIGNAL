package main

import (
	"fmt"
)

func matrixElementsSum(matrix [][]int)int{
	aboveValue := 0
	total := 0
	for idx := range matrix{
		for idx2 := range matrix[idx]{
			if idx == 0{
				continue
			}else{
				aboveValue = matrix[idx -1][idx2]
				if aboveValue == 0{
					matrix [idx][idx2] =0
				}
			}
		}
	}
	fmt.Println(matrix)

	for idx := range matrix{
		for _, val2 := range matrix[idx]{
			total += val2
		}
	}
	return total
}


func main() {
	matrix := [][]int{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}}

	fmt.Printf("Total price for all the rooms suitable for the matrix %v is %d", matrix, matrixElementsSum(matrix))
	}
