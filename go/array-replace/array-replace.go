package main

import(
	"fmt"
)

func arrayReplace(inputArray []int, elemToReplace, substitutionElem int)[]int{
	for idx, val := range inputArray{
		if val == elemToReplace{
			inputArray[idx] = substitutionElem
		}
	}
	return inputArray
}

func main(){
	inputArray := []int{1, 2, 1}
	elemToReplace := 1
	substitutionElem := 3

	fmt.Printf("Subtituing all the %d for %d in the array the answer is: %v", elemToReplace, substitutionElem, arrayReplace(inputArray, elemToReplace, substitutionElem))
}