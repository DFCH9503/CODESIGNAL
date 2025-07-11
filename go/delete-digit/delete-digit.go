package main

import (
	"fmt"
	"strconv"
)

func deleteDigit(n int) int {
	s := strconv.Itoa(n)
	maxNum := 0

	for i := 0; i < len(s); i++ {
		r := s[:i] + s[i+1:]
		num, _ := strconv.Atoi(r)
		if num > maxNum {
			maxNum = num
		}
	}


	return maxNum
}

func main() {
	n := 152

	fmt.Printf("For the integer %d the maximum maximal number you can obtain by deleting exactly one digit from it is: %d", n, deleteDigit(n))
}