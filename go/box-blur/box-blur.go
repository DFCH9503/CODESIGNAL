package main

import (
	"fmt"
	"math"
)

func boxBlur(image [][]int)[][]int{
	res := [][]int{}
	top := 0
	mid := 0
	bot := 0
	for i:= 0; i < len(image) - 2; i++{
		partialMatrix := []int{}
		for j:= 0; j < len(image[i]) - 2; j++{
			top = image[i][j]+image[i][j+1]+image[i][j+2]
			mid = image[i+1][j]+image[i+1][j+1]+image[i+2][j+2]
			bot = image[i+2][j]+image[i+2][j+1]+image[i+2][j+2]
			partialMatrix = append(partialMatrix, int(math.Floor(float64(((top+mid+bot)/9)))))
		}
		res = append(res, partialMatrix)
	} 
	return res
}

func main(){
	image := [][]int{{1, 1, 1}, {1, 7, 1}, {1, 1, 1}}
	// image := [][]int{{7, 4, 0, 1}, {5, 6, 2, 2}, {6, 10, 7, 8}, {1, 4, 2, 0}}
	fmt.Printf("The blurred image is: %v", boxBlur(image))
}