package main

import(
	"fmt"
)

func alternatingSums(a []int)[]int{
	sumEven := 0
    sumOdd := 0

	for idx, val := range a{
		if idx%2 == 0{
            sumEven += val
        }else{
            sumOdd += val
        }
	}
	return []int{sumEven, sumOdd}
}

func main() {
	a := []int{50, 60, 60, 45, 70}
	fmt.Printf("The answer for the array %v is %d", a, alternatingSums(a))
}