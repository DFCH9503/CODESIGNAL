function palindromeRearranging(inputString){
    let splitInput = inputString.split('')
    let numberOfLetters = []
    let keysInNumberOfLetters = []
    let valuesInNumberOfLetters = []
    let counter = 0

    splitInput.map(e => {
        if(numberOfLetters[e]){
            numberOfLetters[e]++
        }else{
            numberOfLetters[e] = 1
        }
    })

    
    keysInNumberOfLetters = Object.keys(numberOfLetters)
    valuesInNumberOfLetters = Object.values(numberOfLetters)
    
    for(i = 0; i < valuesInNumberOfLetters.length; i++){
        if(valuesInNumberOfLetters[i] % 2 !== 0){
            counter++
        }
    }


    return ! (counter > 1)
}

let inputString = "aabb"
console.log(palindromeRearranging(inputString))