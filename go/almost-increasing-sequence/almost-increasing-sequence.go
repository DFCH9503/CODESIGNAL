package main

import(
	"fmt"
)

func almostIncreasingSequence(sequence []int)bool{
	counter := 0
	s := sequence

	for i := 1 ; i < len(s) ; i++{
		if s[i-1] >= s[i]{
			counter ++
			if counter > 1 {
				return false
			}
			if s[i-2] >= s[i] && s[i-1]>= s[i+1]{
				return false
			}
		}
	}
	return true
}

func main() {
	sequence := []int{1 , 3 , 2 , 1}
	fmt.Printf("The array %v is almost increasing: %v", sequence, almostIncreasingSequence(sequence))
}