package main

import(
	"fmt"
)

func longestDigitsPrefix(inputString string)string{
	res := ""
	for _, val := range inputString{
		if val < 48 || val > 57{
			break
		}
		res += string(val)
	}
	return res
}

func main(){
	inputString := "123aa1"

	fmt.Printf("The longest digit prefix in the inputString = %s is: %s ", inputString, longestDigitsPrefix(inputString))

}