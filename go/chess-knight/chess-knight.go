package main

import(
	"fmt"
)

func chessKnight(cell string) int{

	knightX, knightY := int(rune(cell[0])), int(rune(cell[1]))

	possibleKnightX, possibleKnightY := []int{knightX - 2, knightX - 1, knightX + 1, knightX + 2}, []int{knightY - 2, knightY - 1, knightY + 1, knightY + 2}

	for idx, val := range possibleKnightX{
		if val < 97 || val > 104{
			possibleKnightX[idx] = -1
		}
		if possibleKnightY[idx] < 49 || possibleKnightY[idx] > 56{
			possibleKnightY[idx] = -1
		}

	}

	movements := []int{
	possibleKnightX[0] + possibleKnightY[1],
    possibleKnightX[0] + possibleKnightY[2],
    possibleKnightX[1] + possibleKnightY[0],
    possibleKnightX[1] + possibleKnightY[3],
    possibleKnightX[2] + possibleKnightY[0],
    possibleKnightX[2] + possibleKnightY[3],
    possibleKnightX[3] + possibleKnightY[1],
    possibleKnightX[3] + possibleKnightY[2]}
	
	validMovements := 0

	for _, val := range movements{
		if val >= 146 && val <= 160{
			validMovements++
		}
	}

	return validMovements
}

func main(){
	cell := "a1"

	fmt.Printf("For a knight located in the cell %s of a standard chess board there could be %d valid movements", cell, chessKnight(cell))
}