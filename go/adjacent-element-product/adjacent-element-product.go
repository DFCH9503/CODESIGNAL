package main

import(
	"fmt"
)

func adajacentElementProduct(inputArray []int)int{
	res := 0
	adjacentProduct := 0
	for i := 0 ; i < len (inputArray) - 1 ; i++{
		adjacentProduct = inputArray[i] * inputArray[i + 1]
		if adjacentProduct > res {
			res = adjacentProduct
		}
	} 
	return res
}

func main() {
	inputArray := []int{3, 6, -2, -5, 7, 3}
	fmt.Printf("The largest adjacent product in the %v array is %d", inputArray, adajacentElementProduct(inputArray))
}