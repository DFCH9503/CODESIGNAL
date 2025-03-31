function reverseParentheses(s){
    let x=s
    while(x.includes('(')){
        let searchSecondP=x.indexOf(')')
        let searchFirstP=x.lastIndexOf('(',searchSecondP)
        let stringToReverse=x.slice((searchFirstP+1),searchSecondP)
        let reversedString= Array.from(stringToReverse).reverse().join('')
        
        x=x.slice(0,searchFirstP)+reversedString+x.slice(searchSecondP+1)

    }
    return x
}

let s = "a(bc)de"
console.log(reverseParentheses(s))