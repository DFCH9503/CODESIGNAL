function boxBlur(image){
    const result=[]
    for(let i=0;i<image.length-2;i++){
        const partialMatrix=[]
        for(let j=0;j<image[i].length-2;j++){
            const top=image[i][j]+image[i][j+1]+image[i][j+2]
            const mid=image[i+1][j]+image[i+1][j+1]+image[i+1][j+2]
            const bot=image[i+2][j]+image[i+2][j+1]+image[i+2][j+2]
            partialMatrix.push(Math.floor(((top+mid+bot)/9))) 
    }
    result.push(partialMatrix)
    }

return result
}

let image = [[1, 1, 1], [1, 7, 1], [1, 1, 1]]
console.log(boxBlur(image))