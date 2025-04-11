package main

import(
	"fmt"
)

func palindromeRearranging(inputString string) bool{
	mp := make(map [byte]int)
	for i := 0; i < len(inputString); i++{
		if _, exists := mp[inputString[i]]; exists{
			mp[inputString[i]] ++
		}else{
			mp[inputString[i]] = 1
		}
	}

	counterOneLetter := 0
	for _, value := range mp{
		if value % 2 != 0{
			counterOneLetter++
		}
	}
	return !(counterOneLetter>1)
}

func main(){
	inputString := "aabb"
	fmt.Printf("The answer is: %v", palindromeRearranging(inputString))
}