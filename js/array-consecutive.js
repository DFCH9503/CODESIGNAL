function makeArrayConsecutive2(statues){
    let maxArray = Math.max(...statues)
    let minArray = Math.min(...statues)
    let diffArray = maxArray-minArray
    let result = (diffArray - statues.length) + 1
    return result

}

let statues = [6 , 2 , 3 , 8]
console.log(makeArrayConsecutive2(statues))
