package main

import (
	"fmt"
	"strconv"
)

func isLucky(n int)bool{
	nString := strconv.Itoa(n)
	middle := len(nString)/2
	resLeft := 0
	resRight := 0

	for i :=0 ; i < middle ; i++{
		resLeft += int(nString[i])
		resRight += int(nString[len(nString)-1-i])
	}

	return resLeft == resRight
}

func main() {
	n := 239017
	fmt.Printf("The ticket: %d is %v", n, isLucky(n))
}