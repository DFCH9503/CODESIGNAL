function alternatingSums(a){
    let sumEven = 0
    let sumOdd = 0

    for (let i=0 ; i<a.length ; i++){
        if (i%2 == 0){
            sumEven += a[i]
        }else{
            sumOdd += a[i]
        }
    }
    return [sumEven, sumOdd]
}

let a = [50, 60, 60, 45, 70]
console.log(alternatingSums(a))