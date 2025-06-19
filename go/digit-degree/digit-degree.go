package main

import (
	"fmt"
	"strconv"
)

func sumDigits(n int) int{
	nStr := strconv.Itoa(n)
	nSum := 0
	for _, val := range nStr{
		digit, _ := strconv.Atoi(string(val))
		nSum += digit
	}
	return nSum
}

func digitDegree(n int) int {
	if n < 10{
        return 0
	}else{
		return 1 + digitDegree(sumDigits(n))
	}
	
}

func main(){
	n := 91

	fmt.Printf("The digit degree of the number (number of times we need to replace this number with the sum of its digits until we get to a one digit number) %d is: %d", n,  digitDegree(n))
}