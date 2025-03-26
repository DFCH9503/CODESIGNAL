function allLongestString(inputArray){
    let maxLetterCount=0
    let letterCount=[]
    let result=[]
    for(i=0;i<inputArray.length;i++){
        letterCount[i]=inputArray[i].length
    }

    maxLetterCount=Math.max(...letterCount)

    result=inputArray.filter(x=>x.length==maxLetterCount)

    return result
}

let inputArray = ["aba", "aa", "ad", "vcd", "aba"]
console.log(allLongestString(inputArray))