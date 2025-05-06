function arrayReplace(inputArray, elemToReplace, substitutionElem){
    let res = inputArray.map(e => e == elemToReplace ? substitutionElem : e)
    return res
}

let inputArray = [1, 2, 1], elemToReplace = 1, substitutionElem = 3  

console.log(arrayReplace(inputArray, elemToReplace, substitutionElem))