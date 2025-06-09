function avoidObstacles(inputArray){
    for(let i = 1; ; i++){
    if(inputArray.every(e => e % i != 0)){
        return i
    }
}
}

let inputArray = [5, 3, 6, 7, 9]
console.log(avoidObstacles(inputArray))