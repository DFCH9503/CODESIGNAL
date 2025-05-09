package main

import(
	"fmt"
)

func minesweeper(matrix [][]bool)[][]int{
	rows := len(matrix)
	if rows == 0 {
		return [][]int{}
	}
	cols := len(matrix[0])

	solBoard := make([][]int, rows)
	for i := range solBoard {
		solBoard[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// Above
			if i > 0 && matrix[i-1][j] {
				solBoard[i][j]++
			}
			// Above-Right
			if i > 0 && j+1 < cols && matrix[i-1][j+1] {
				solBoard[i][j]++
			}
			// Right
			if j+1 < cols && matrix[i][j+1] {
				solBoard[i][j]++
			}
			// Bottom-Right
			if i+1 < rows && j+1 < cols && matrix[i+1][j+1] {
				solBoard[i][j]++
			}
			// Bottom
			if i+1 < rows && matrix[i+1][j] {
				solBoard[i][j]++
			}
			// Bottom-Left
			if i+1 < rows && j > 0 && matrix[i+1][j-1] {
				solBoard[i][j]++
			}
			// Left
			if j > 0 && matrix[i][j-1] {
				solBoard[i][j]++
			}
			// Top-Left
			if i > 0 && j > 0 && matrix[i-1][j-1] {
				solBoard[i][j]++
			}
		}
	}

	return solBoard
}

func main(){
	matrix := [][]bool{{true, false, false}, {false, true, false}, {false, false, false}}
	fmt.Printf("the minesweeper game represented by the matrix %v could be represented (by total number of mines in the neighboring cells) with the matrix %v", matrix, minesweeper(matrix))
}