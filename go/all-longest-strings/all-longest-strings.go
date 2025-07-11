package main

import(
	"fmt"
)

func allLongestStrings(inputArray []string)[]string{
	maxLetterCount := 0
	var res []string
	for _, val := range inputArray{
		if len(val) > maxLetterCount{
			maxLetterCount = len (val)
		}
	}

	for _, val := range inputArray{
		if len(val) == maxLetterCount{
			res = append(res, val)
		}
	}
	return res
}

func main() {
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}
	fmt.Printf("The answer for the array %v is %v", inputArray, allLongestStrings(inputArray))
}