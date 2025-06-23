function isBeautifulString(inputString){
    const lettersCharCode = Array.from(Array(26)).map((e, i) => i + 97)
    const alphabet = lettersCharCode.map((x) => String.fromCharCode(x))
    const letters = Array(alphabet.length).fill(0)
    
    for (let letter of inputString) {
        letters[alphabet.indexOf(letter)]++
    }

    
    for (let i = 0; i < letters.length; i++) {
        if (letters[i + 1] > letters[i]) return false
    }
    return true

}

let inputString = "bbbaacdafe"

console.log(`The string "${inputString}" is beautiful:`, isBeautifulString(inputString))