function longestWord(text){
    let pattern = /[^a-z^A-Z]/
    
    text = text.split(pattern)
    
    let longest = text.reduce((a, b) => {
        return a.length > b.length ? a : b
    })
    
    return longest
}

let text = "Ready, steady, go!"

console.log(`The longest word in the text '${text}' is:`, longestWord(text))