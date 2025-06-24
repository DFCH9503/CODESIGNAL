function differentSymbolsNaive(s){
    // let string = String.prototype.concat.call(...new Set(s))

    let string = s.split("")
    let stringSet = new Set(string)
    return stringSet.size
}

let s = "cabca"

console.log(`The number of different characters in the string "${s}" is:`, differentSymbolsNaive(s))