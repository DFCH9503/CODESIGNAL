package main

import(
	"fmt"
	"slices"
)

func makeArrayConsecutive2(statues []int)int{
	max := slices.Max(statues)
	min := slices.Min(statues)
	diff := max - min
	return diff - len(statues) + 1
}

func main() {
	statues := []int{6 , 2 , 3 , 8}
	fmt.Printf("The answer for the array %v is %d", statues, makeArrayConsecutive2(statues))
}