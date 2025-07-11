package main

import (
	"fmt"
	"strings"

)

func reverseString(s string)string{
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseParentheses(s string)string{
	x := s

	for strings.Contains(x,"("){
		searchSecondP := strings.Index(x, ")")
		searchFirstP := strings.LastIndex(x[:searchSecondP], "(")
		stringToReverse := x[searchFirstP+1 : searchSecondP]
		reversedString := reverseString(stringToReverse)
		x = x[:searchFirstP] + reversedString + x[searchSecondP+1:] 
	}
	return x
}

func main() {
	s := "a(bc)de"
	fmt.Printf("The answer is: %v", reverseParentheses(s))
}