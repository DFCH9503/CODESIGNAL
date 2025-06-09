function adajacentElementProduct(inputArray){

    let product = 1
    let arrayProducts = []

    
    for(let i = 0; i < inputArray.length-1; i++){
        product = inputArray[i] * inputArray[i+1]
        arrayProducts.push(product)
    }
    
    
    return maxProductAdjacent = Math.max(...arrayProducts)
}

let  inputArray = [3, 6, -2, -5, 7, 3]
adajacentElementProduct(inputArray)
