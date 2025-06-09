function arrayMaxConsecutiveSum(inputArray, k) {
    let max = inputArray.slice(0 , k).reduce((a,b) => a + b)
        
    let cur = max
    for(let i = k; i < inputArray.length; i++) {
        cur = cur + inputArray[i] - inputArray[i - k]
        if(cur > max)
            max = cur
    }
    return max
}

let inputArray = [2, 3, 5, 1, 6], k = 2

console.log(`The maximal possible sum of some of its k = ${k} consecutive elements of the array [${inputArray}] is:`, arrayMaxConsecutiveSum(inputArray, k))