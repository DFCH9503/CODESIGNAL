function arrayChange(inputArray){
    let increasingArray = []
    let originalArray = inputArray.slice(0)

    for(i = 0;i < inputArray.length; i++){
        while(inputArray[i] >= inputArray[i + 1]){
            inputArray[i + 1] += 1
        }
    }

    for(i = 0; i < inputArray.length; i++){
        increasingArray.push(inputArray[i] - originalArray[i])
    }
    return increasingArray.reduce((acc,curr) => acc + curr, 0)
}


let inputArray = [1, 1, 1]
console.log(arrayChange(inputArray))