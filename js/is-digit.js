function isDigit(symbol){

    //without reg exp
    let arrayDigit = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    let arrayDigitString = arrayDigit.map(e => String(e))
    
    return arrayDigitString.includes(symbol)

    //with regexp
    // return /^\d{1}$/.test(symbol)
}

let symbol = "0"

console.log(`The symbol ${symbol} is a valid digit:`, isDigit(symbol))