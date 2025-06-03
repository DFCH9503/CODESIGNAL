package main

import(
	"fmt"
	"slices"
)

func firstDigit(inputString string)string{
	start, end := 48 , 57
	numbers := make([]int, 0)
    for i := start; i <= end; i++ {
        numbers = append(numbers, i)
    }
	for _, val := range inputString{
		if slices.Contains(numbers, int(val)){
			return string(rune(val))
		} 
	}
	return "0"
}

func main(){
	inputString := "var_1__Int"
	fmt.Printf("The leftmost digit that occurs in %s is: %s \n", inputString, firstDigit(inputString))
}