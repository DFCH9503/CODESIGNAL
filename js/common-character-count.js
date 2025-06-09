function commonCharacterCount(s1, s2){

    let commonCharacters = 0
    let s1Split = s1.split('')
    let s2Split = s2.split('')

    s1Split.forEach(element => {
        if(s2Split.includes(element)){
            commonCharacters ++
            s2Split.splice(s2Split.indexOf(element), 1)
        }
        
    })

    return commonCharacters
}

let s1="aabcc", s2 ="adcaa"
console.log(commonCharacterCount(s1, s2))