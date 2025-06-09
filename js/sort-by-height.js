function sortByHeight(a){
    let arrayPerson = []

    for(i = 0;i < a.length; i++){
        if(a[i] != -1){
            arrayPerson.push(a[i])
        }
    }
    arrayPerson.sort((a,b) => a - b)

    for(let i = 0; i < a.length; i++) {
        if(a[i] != -1) {
            a[i] = arrayPerson.shift()
        }
    }
    
    return a
}

let a = [-1, 150, 190, 170, -1, -1, 160, 180]
console.log(sortByHeight(a))