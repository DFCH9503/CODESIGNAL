function circleOfNumbers(n ,firstNumber){
    return (n / 2 + firstNumber) % n
}

let n = 10, firstNumber = 2

console.log(`The radially opposite position to ${firstNumber} in a circle starting with 0 and ending with ${n-1} is:`, circleOfNumbers(n, firstNumber))