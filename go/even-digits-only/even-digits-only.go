package main

import(
	"fmt"
	"strconv"
)

func evenDigitsOnly(n int)bool{
	nString := strconv.Itoa(n)

	for _, val := range nString{
		if val % 2 != 0{
			return false
		} 
	}
	return true
}

func main(){
	n := 642386
	fmt.Printf("All the digits in the integer %d are even: %v", n, evenDigitsOnly(n))
}