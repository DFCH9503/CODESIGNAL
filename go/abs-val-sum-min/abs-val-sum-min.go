package main

import(
	"fmt"
	"math"
)

func absoluteValuesSumMinimization(a []int)int{
	if len(a) % 2 == 0{
		return a[len(a) / 2 -1]
	}else{
		idx := int(math.Floor(float64(len(a) / 2)))
		return  a[idx]
	}
}

func main(){
	a := []int{2, 4, 7}
	fmt.Printf("The smallest integer from the sorted array %v that mimizes the sum abs(a[0] - x) + abs(a[1] - x) + ... + abs(a[a.length - 1] - x) is %d", a,  absoluteValuesSumMinimization(a))
}