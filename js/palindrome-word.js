//brute approach


function isPalindromeWord(inputString){

    let arrayString = inputString.split("")
    
    let trueArray = []
    
    for(let i = 0;i <= arrayString.length-1; i++){
        let firstLetter = arrayString[i]
        let lastLetter = arrayString[arrayString.length - ( i + 1)]
        if(firstLetter == lastLetter){
            trueArray.push('true')
        }else{
            trueArray.push('false')
        }
        
    }
    
    if(trueArray.indexOf('false') == -1){
        return 'palindrome'
    }else{
        return 'not palindrome'
    }
}

let s = "aobuboa"

isPalindromeWord(s)