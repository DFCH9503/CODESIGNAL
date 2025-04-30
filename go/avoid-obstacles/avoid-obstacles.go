package main

import(
	"fmt"
)


func avoidObstacles(inputArray []int)int{
	for i := 1; ; i++{
		isValid := true
		for _, val := range inputArray{
			if val % i == 0{
				isValid = false
				break
			}
		}
		if isValid{
			return i
		}
	}
}

func main(){
	inputArray := []int{5, 3, 6, 7, 9}
	fmt.Printf("the minimal length of the jump enough to avoid all the obstacles represented by %v is: %d", inputArray, avoidObstacles(inputArray))
}