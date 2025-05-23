package main

import (
	"fmt"
)

func chessBoardCellColor(cell1, cell2 string)bool{
	counter1, counter2 := 0, 0
	for idx, val := range cell1{
		counter1 += int(val)
		counter2 += int(cell2[idx])
	}
	return counter1 % 2 == 0 && counter2 % 2 == 0 
}

func main(){
	cell1, cell2 := "A1",  "C3"

	fmt.Printf("The cells %s and %s of a standard chess board have the same color: %v", cell1, cell2, chessBoardCellColor(cell1, cell2))
}