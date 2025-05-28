package main

import(
	"fmt"
)

func extractEachKth(inputArray []int, k int)[]int{
	res := []int{}
	for idx, val := range inputArray{
		if (idx + 1) % k != 0{
			res = append(res, val)
		}
	}
	return res
}

func main(){
	inputArray, k := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3

	fmt.Printf("If you remove each %d element from the input array %v you'll get %v", k, inputArray, extractEachKth(inputArray, k))
}