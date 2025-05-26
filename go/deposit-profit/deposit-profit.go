package main

import(
	"fmt"
)

func depositProfit(deposit, rate, threshold int)int{
	n := 0
    for deposit < threshold{
        n++
        deposit += deposit * rate / 100
    }
    
    return n
}

func main(){
	deposit, rate, threshold := 100, 20, 170

	fmt.Printf("To pass the threshold %d with an initial deposit of %d and with an annual rate of %d%% you'll need %d years", threshold, deposit, rate, depositProfit(deposit, rate, threshold))
}