package main

import(
	"fmt"
)

func arrayMaximalAdjacentDifference(inputArray []int)int{
	res := 0
	diff := 0
	for i := 0; i < len(inputArray) - 1; i++{
		diff = inputArray[i] - inputArray[i + 1]
		if diff < 0 {
			diff *= -1
		}
		if res < diff{
			res = diff
		}
	}
	return res
}

func main(){
	inputArray := []int{2, 4, 1, 0}
	fmt.Printf("The max diff between adjacent elements for the array %v is %d", inputArray, arrayMaximalAdjacentDifference(inputArray))
}