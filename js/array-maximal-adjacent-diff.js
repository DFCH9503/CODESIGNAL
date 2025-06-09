function arrayMaximalAdjacentDifference(inputArray){
    let diffArray = []
    let maxDiff = 0
    
    for(i = 0;i < inputArray.length -1; i++){
        diffArray.push(inputArray[i] - inputArray[i + 1])
        if(diffArray[i] < 0){
            diffArray[i] *= -1
        }
    }
    
    return maxDiff=Math.max(...diffArray)
}

let inputArray = [2, 4, 1, 0]
console.log(arrayMaximalAdjacentDifference(inputArray))