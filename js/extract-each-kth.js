function extractEachKth(inputArray, k){
    return inputArray.filter((e, i) => {
        return (i + 1) % k != 0
    })
}

let inputArray = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3

console.log(`If you remove each ${k} element from the input array [${inputArray}] you'll get`,extractEachKth(inputArray, k))