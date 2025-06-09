function chessBoardCellColor(cell1, cell2){
    const alpha = Array.from(Array(8)).map((e, i) => i + 65);
    const letters = alpha.map((x) => String.fromCharCode(x));
    const numbers = Array.from(Array(8)).map((e, i) => i + 1);

    let blacks = []

    for(let i = 0;i < numbers.length; i++){
        for(let j = 0; j < letters.length; j++){
            if(i % 2 == 0 && j % 2 == 0){
                blacks.push(letters[j] + numbers[i])
            }else if(i % 2 == 1 && j % 2 == 1){
                blacks.push(letters[j] + numbers[i])     
            }
        }
    }

    let isCell1 = blacks.includes(cell1)
    let isCell2 = blacks.includes(cell2)


    return isCell1 && isCell2

}

let cell1 = "A1", cell2 = "C3"

console.log(`The cells ${cell1} and ${cell2} of a standard chess board have the same color: `, chessBoardCellColor(cell1, cell2) )