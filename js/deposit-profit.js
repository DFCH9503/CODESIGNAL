function depositProfit(deposit, rate, threshold){
    let n = 0
    while(deposit < threshold){
        n++
        deposit += deposit * rate / 100
    }
    
    return n
}

let deposit = 100, rate = 20, threshold = 170

console.log(`To pass the threshold ${threshold} with an initial deposit of ${deposit} and with an annual rate of ${rate}% you'll need`, depositProfit(deposit, rate, threshold), `years`)