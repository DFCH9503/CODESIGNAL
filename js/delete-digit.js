function deleteDigit(n){
    let s = String(n).split('').map(Number)
    return Math.max(...s.map((el, i) => {
        let r = s.slice(); 
        r.splice(i, 1);
        return Number(r.join(''))
        }))
}

let n = 152

console.log(`For the integer ${n} the maximum maximal number you can obtain by deleting exactly one digit from it is:`, deleteDigit(n))