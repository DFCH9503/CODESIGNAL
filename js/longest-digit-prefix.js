function longestDigitsPrefix(inputString){
    let res = ""
    for (const element of inputString) {
        if(!/\d/.test(element))
            break
        res += element
    }   
    return res
}

let inputString = "123aa1"

console.log(`The longest digit prefix in the inputString = ${inputString} is: `, longestDigitsPrefix(inputString))