package main

import(
	"fmt"
)

func isPalindromeWord(s string)bool{
	left := 0
	rigth := len(s)-1

	for idx := range s{
		if s[left] != s [rigth]{
			return false
		}
		left += idx
		rigth -= idx
	}
	return true
}

func main() {
	s := "asbuboa"
	fmt.Printf("The word %s is palindrome: %t", s, isPalindromeWord(s))
}