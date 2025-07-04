package main

import(
	"fmt"
)

func bishopAndPawn(bishop, pawn string) bool{
	bishopX, bishopY := rune(bishop[0]), rune(bishop[1])
	pawnX, pawnY := rune(pawn[0]), rune(pawn[1])
	
	if bishopX + pawnY == bishopY + pawnX || bishopX + bishopY == pawnX + pawnY {
    	return true
	}

return false
}

func main(){
	bishop, pawn := "a1", "c3"

	fmt.Printf("A bishop placed in the square %s of a standard chess board  can capture a pawn placed in the square %s: %v", bishop, pawn,  bishopAndPawn(bishop, pawn))
}