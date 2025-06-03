function firstDigit(inputString){

    //wITHOUT REGULAR EXPRESSIONS
    // let numbers = Array.from(Array(10)).map((e, i) => i + 48).map((x) => String.fromCharCode(x))
    // let stringSplit = inputString.split('')
    // for(let i = 0; i < stringSplit.length; i++){
    //     if(numbers.includes(stringSplit[i])){
    //         return stringSplit[i]
    //     }
    // }

    //WITH REGULAR EXPRESSIONS
    return inputString[inputString.search(/\d/)]
}

let inputString = "var_1__Int"
console.log(`The leftmost digit that occurs in ${inputString} is:`, firstDigit(inputString))