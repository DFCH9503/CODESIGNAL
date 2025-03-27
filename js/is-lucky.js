function isLucky(n){

    let middle = n.toString().length/2

    let nRight = String(n).split('').map(Number).reduce((a,b)=>a+b)/2
    let nLeft = String(n).split('').slice(0,middle).map(Number).reduce((a,b)=>a+b)

    return nRight == nLeft

}

let n=239017
console.log(isLucky(n))