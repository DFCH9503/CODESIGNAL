function addBorder(picture){
    let character='*'
    let width=picture[0].length
    let topBottomBorder=character.repeat(width+2)

    let partialArray=[]

    for(i=0;i<picture.length;i++){
        partialArray.push(character+picture[i]+character)
    }

    partialArray.unshift(topBottomBorder)
    partialArray.push(topBottomBorder)

    return partialArray
}

let picture = ["abc", "ded"]
console.log(addBorder(picture))