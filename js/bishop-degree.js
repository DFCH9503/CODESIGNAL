function bishopAndPawn(bishop, pawn){
let board={
    'a': 1,
    'b': 2,
    'c': 3,
    'd': 4,
    'e': 5,
    'f': 6,
    'g': 7,
    'h': 8   
}

let bishopX = board[bishop[0]]
let bishopY = Number(bishop[1])

let pawnX = board[pawn[0]]
let pawnY = Number(pawn[1])


if(bishopX + pawnY == bishopY + pawnX || bishopX + bishopY == pawnX + pawnY){
    return true
}

return false
}

let bishop = "a1", pawn = "c3"

console.log(`A bishop placed in the square ${bishop} of a standard chess board  can capture a pawn placed in the square ${pawn}: `, bishopAndPawn(bishop, pawn))