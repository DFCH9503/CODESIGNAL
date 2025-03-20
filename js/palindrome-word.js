//brute approach

let arrayString=inputString.split("")

    let trueArray=[]

    for(i=0;i<=arrayString.length-1;i++){
        firstLetter=arrayString[i]
        lastLetter=arrayString[arrayString.length-(i+1)]
        if(firstLetter==lastLetter){
            trueArray.push('true')
        }else{
            trueArray.push('false')
        }
        
    }
    console.log(trueArray)
    if(trueArray.indexOf('false')==-1){
        return console.log('palindrome')
    }else{
        return console.log('not palindrome')
    }