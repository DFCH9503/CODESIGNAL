function minesweeper(matrix) {

    const solBoard=[]
    for(let i=0; i<matrix.length;i++){
        solBoard.push([])
        for(let j=0; j<matrix[i].length;j++){
            solBoard[i][j]=0
            //Above
            if(matrix[i-1]!==undefined){
                if(matrix[i-1][j]){
                    solBoard[i][j]++
                }
            }
            //Above-Right
            if(matrix[i-1]!==undefined){
                if(matrix[i-1][j+1]){
                    solBoard[i][j]++
                }
            }
            //Right
            if(matrix[i][j+1]!==undefined){
                if(matrix[i][j+1]){
                    solBoard[i][j]++
                }
            }
            //Bottom-Right
            if(matrix[i+1]!==undefined){
                if(matrix[i+1][j+1]){
                    solBoard[i][j]++
                }
            }
            //Bottom
            if(matrix[i+1]!==undefined){
                if(matrix[i+1][j]){
                    solBoard[i][j]++
                }
            }
            //Bottom-Left
            if(matrix[i+1]!==undefined){
                if(matrix[i+1][j-1]){
                    solBoard[i][j]++
                }
            }
            //Left
            if(matrix[i][j-1]!==undefined){
                if(matrix[i][j-1]){
                    solBoard[i][j]++
                }
            }
            //Top-Left
            if(matrix[i-1]!==undefined){
                if(matrix[i-1][j-1]){
                    solBoard[i][j]++
                }
            }
            
            
            
        }
    }
    return solBoard

}

let matrix = [[true, false, false], [false, true, false], [false, false, false]]

console.log(minesweeper(matrix))