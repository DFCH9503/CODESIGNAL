package main

import(
	"fmt"
)


func max(a, b int) int{
	if a < b{
		return b
	}
	return a
}
func knapsackLight(value1, weight1, value2, weight2, maxW int) int{
	if maxW >= (weight1 + weight2){
        return value1 + value2
    }else if maxW >= weight1 && maxW >= weight2{
		return max(value1, value2)
    }else if maxW >= weight1 && maxW < weight2{
        return value1
    }else if maxW >= weight2 && maxW < weight1{
        return value2
    }else{
        return 0
    }
}

func main(){
	value1 := 10
	weight1 := 5
	value2 := 6
	weight2 := 4
	maxW := 8

	fmt.Printf("The total max value that you can take assuming that your max weight capacity is maxW = %d and you can’t come back is: %d", maxW,  knapsackLight(value1, weight1, value2, weight2, maxW))
}