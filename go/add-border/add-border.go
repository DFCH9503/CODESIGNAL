package main

import (
	"fmt"
	"strings"
)

func addBorder(picture []string)[]string{
	character := "*"
	width := len(picture[0])
	topBottomBorder := strings.Repeat(character, width+2)

	res := []string{}


	for i := 0 ; i < len(picture) ; i++{
		res = append(res, character+picture[i]+character)
	} 

	res = append([]string{topBottomBorder}, res...)
	res = append(res, topBottomBorder)
	return res
}

func main() {
	picture := []string{"abc", "ded"}
	fmt.Printf("The answer for the array %v is %v", picture, addBorder(picture))
}