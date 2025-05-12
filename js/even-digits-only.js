function evenDigitsOnly(n){
    return String(n).split('').every(e => e % 2 == 0)
}

let n = 248622

console.log(evenDigitsOnly(n))