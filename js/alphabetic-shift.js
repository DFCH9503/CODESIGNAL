function alphabeticShift(inputString){
    const alpha = Array.from(Array(26)).map((e, i) => i + 97);
    const alphabet = alpha.map((x) => String.fromCharCode(x));
    const inputStringSplit = inputString.split('')
    let charChange = []
    let arrayNewChar = []
    for(let i = 0; i < inputStringSplit.length; i++){
        charChange.push(alphabet.indexOf(inputStringSplit[i]) + 1)
    }
    for(let i = 0; i < charChange.length; i++){
        if(charChange[i] == 26){
            arrayNewChar.push('a')
        }else{
            arrayNewChar.push(alphabet.at(charChange[i]))
        }
    }
    return arrayNewChar.join('')
}

let inputString = "crazy"

console.log(alphabeticShift(inputString))