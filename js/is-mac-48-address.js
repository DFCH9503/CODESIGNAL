function isMAC48Address(inputString){
    let regExp = /^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/

    return inputString.match(regExp) ? true : false
}

let inputString = "00-1B-63-84-45-E6"

console.log(`The string ${inputString} is a valid standard IEEE 802 MAC-48 address:`, isMAC48Address(inputString))