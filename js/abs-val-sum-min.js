function absoluteValuesSumMinimization(a){
    return a.length % 2 == 0 ? a[a.length / 2 - 1] : a[Math.floor(a.length / 2)]
}

let a = [2, 4, 7]

console.log(`The smallest integer from the sorted array ${a} that mimizes the sum abs(a[0] - x) + abs(a[1] - x) + ... + abs(a[a.length - 1] - x) is`, absoluteValuesSumMinimization(a))