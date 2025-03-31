package main

import (
	"fmt"
	"slices"
	
)

func sortByHeight(a []int)[]int{
	slicePerson := []int{}
	for _, val := range a{
		if val != -1{
			slicePerson = append(slicePerson, val)
		}
	}
	slices.Sort(slicePerson)

	for idx, val := range a{
		if val !=-1{
			a[idx] = slicePerson[0] 
			slicePerson = slicePerson[1:] //the two lines are like a shift in js
		}
	}
	return a
}

func main() {
	a := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	fmt.Printf("The answer is: %v", sortByHeight(a))
}