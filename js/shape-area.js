function shapeArea(n){
    if(n == 1){
            area = 1
        }else{
            area = shapeArea(n - 1) + ((n - 1) * 4)
        }
    
        return area

}

let n = 2
console.log(shapeArea(n))
