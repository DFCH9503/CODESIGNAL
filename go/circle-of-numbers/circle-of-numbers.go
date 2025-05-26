package main

import(
	"fmt"
)

func circleOfNumbers(n, firstNumber int)int{
	return (n / 2 + firstNumber) % n
}


func main(){
	n, firstNumber := 10 , 2

	fmt.Printf("The radially opposite position to %d in a circle starting with 0 and ending with %d is: %d ",firstNumber, n-1, circleOfNumbers(n, firstNumber))
}
