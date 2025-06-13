package main

import(
	"fmt"
)

func reduce(input []int, initialValue int)int{
	accumulator := initialValue
    for _, value := range input {
        accumulator += value 
    }
    return accumulator
}

func arrayMaxConsecutiveSum(inputArray []int, k int)int{
	max := reduce(inputArray[0 : k], 0)
	curr := max

	for i := k ; i < len(inputArray); i++{
		curr = curr + inputArray[i] - inputArray[i - k]
		if curr > max {
			max = curr
		}
	}

	return max
} 

func main(){
	inputArray, k := []int{2, 3, 5, 1, 6}, 2
	fmt.Printf("The maximal possible sum of some of its k = %d consecutive elements of the array %v is: %d", k, inputArray, arrayMaxConsecutiveSum(inputArray, k))
}