package main

import(
	"fmt"
)

func alphabeticshift(inputString string)string{
	runeSlice := []rune{}
	for _, char := range inputString{
		if char == 122{
			char = 97
			runeSlice = append(runeSlice, rune(int(char)))
		}else{
			runeSlice = append(runeSlice, rune(int(char) + 1))
		}
	}
	return string(runeSlice)
}

func main(){
	inputString := "crazy"

	fmt.Printf("Replacing each letter in the input string %s by the next one in the english alphabet the output is: %s", inputString, alphabeticshift(inputString))	
}