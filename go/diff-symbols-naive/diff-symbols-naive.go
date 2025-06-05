package main

import(
	"fmt"
)

func differentSymbolsNaive(s string)int {
	map1 := make(map[rune]int)

	for _, char := range s{
		map1[char] += 1
	}

	return len(map1)
}

func main(){
	s := "cabca"

	fmt.Printf("The number of different characters in the string %s is: %d", s,  differentSymbolsNaive(s))
}