function variableName(name){
    let regExp = /^[a-zA-Z_]+[a-zA-Z0-9_]*$/
    const found = regExp.test(name);
    return found
}

let name = "qq-q"

console.log(variableName(name))