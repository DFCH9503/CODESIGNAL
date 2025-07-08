function chessKnight(cell){
    let board = { 
    'a': 1,
    'b': 2,
    'c': 3,
    'd': 4,
    'e': 5,
    'f': 6,
    'g': 7,
    'h': 8   
    }
    
    let knightX = board[cell[0]]
    let knightY = Number(cell[1])
    
    let possibleKnightX = [knightX - 2, knightX - 1, knightX + 1, knightX + 2] 
    let possibleKnightY = [knightY - 2, knightY - 1, knightY + 1, knightY + 2] 
    
    for(let i = 0 ; i < possibleKnightX.length; i++){
        if(possibleKnightX[i] <= 0 || possibleKnightX[i] > 8) possibleKnightX[i] = NaN
        
        if(possibleKnightY[i] <= 0 || possibleKnightY[i] > 8) possibleKnightY[i] = NaN
    }
        
    let movements = [
    possibleKnightX[0] + possibleKnightY[1],
    possibleKnightX[0] + possibleKnightY[2],
    possibleKnightX[1] + possibleKnightY[0],
    possibleKnightX[1] + possibleKnightY[3],
    possibleKnightX[2] + possibleKnightY[0],
    possibleKnightX[2] + possibleKnightY[3],
    possibleKnightX[3] + possibleKnightY[1],
    possibleKnightX[3] + possibleKnightY[2]
    ]
    
    let actualMovements = movements.filter(e => !isNaN(e))
    
    return actualMovements.length
}

let cell = "a1"

console.log(`For a knight located in the cell ${cell} of a standard chess board there could be`, chessKnight(cell), `valid movements`)