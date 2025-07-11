package main

import(
	"fmt"
)

func shapeArea(n int)int{
	area := 0
	if n == 1{
		area = 1
	}else{
		area = shapeArea(n - 1)+((n - 1) * 4)
	}

	return area
}

func main() {
	n := 2
	fmt.Printf("The area for n = %d is %d", n, shapeArea(n))
}