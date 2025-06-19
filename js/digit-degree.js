function digitDegree(n){
    if (n < 10){
        return 0
    }else{
        return 1 + digitDegree(Number(String(n).split('').map(Number).reduce((a, b) => a + b)))
    }
}

let n = 91
console.log(`The digit degree of the number (number of times we need to replace this number with the sum of its digits until we get to a one digit number) ${n} is:`, digitDegree(n))