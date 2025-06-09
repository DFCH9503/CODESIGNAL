function matrixElementSum(matrix){
    let flatArray = []
    let total = 0


    let aboveValue = 0
    for(i= 0;i < matrix.length; i++){
        for(j = 0;j < matrix[i].length; j++){
            if(i == 0){
                continue
            }else{
                aboveValue = matrix[i-1][j]
                if(aboveValue == 0){
                    matrix[i][j] = 0
                }
            }
        }
    }

    flatArray = matrix.flat()

    total = flatArray.reduce((acc,curr) => acc + curr, 0)
    return total
}

let matrix = [[0, 1, 1, 2], [0, 5, 0, 0], [2, 0, 3, 3]]
console.log(matrixElementSum(matrix))