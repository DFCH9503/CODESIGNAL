function adajacentElementProduct(inputArray){

    let product=1
    let arrayProducts=[]
    let maxProductAdjacent=0
    
    for(let i=0;i<inputArray.length-1;i++){
        product=inputArray[i]*inputArray[i+1]
        arrayProducts.push(product)
    }
    
    
    return console.log(maxProductAdjacent=Math.max(...arrayProducts))
}

let  inputArray = [3, 6, -2, -5, 7, 3]
adajacentElementProduct(inputArray)
