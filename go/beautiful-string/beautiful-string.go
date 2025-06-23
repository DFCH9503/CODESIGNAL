package main

import(
	"fmt"
)

func isBeautifulString(inputString string) bool{
	m := make(map[rune]int)
	for _, val := range inputString{
		m[val]++
	}
	
	for key := range m{
		if m[key] < m[key+1]{
			return false
		}
	}
	return true
}

func main(){
	inputString := "bbbaacdafe"

	fmt.Printf("The string %s is beautiful: %v", inputString, isBeautifulString(inputString))
}