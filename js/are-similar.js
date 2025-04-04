function areSimilar(a, b){
    let arrayCompairingA=[]
    let arrayCompairingB=[]
    for(i=0;i<a.length;i++){
        if(a[i]!==b[i]){
            arrayCompairingA.push(a[i])
            arrayCompairingB.push(b[i])
        }
    }

    if(arrayCompairingA.join('')===arrayCompairingB.reverse().join('')){
        if(arrayCompairingA.length>2){
            return false
        }else{
            return true
        }
    }else{
        return false
    }
}

let a = [1, 2, 3], b = [2, 1, 3]
console.log(areSimilar(a, b))