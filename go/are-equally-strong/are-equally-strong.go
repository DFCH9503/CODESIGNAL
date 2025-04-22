package main

import (
	"fmt"
	"slices"
)

func areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight int) bool {
	myArms := []int{yourLeft, yourRight}
	friendArms := []int{friendsLeft, friendsRight}

	return slices.Max(myArms) == slices.Max(friendArms) && slices.Min(myArms) == slices.Min(friendArms) 

}

func main() {
	yourLeft := 10
	yourRight := 15
	friendsLeft := 15
	friendsRight := 10

	fmt.Printf("The answer is: %v", areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight))
}