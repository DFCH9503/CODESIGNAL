package main

import(
	"fmt"
	// "slices" //without regexp
	"regexp" //to use it with regexp
)

func isDigit(symbol string) bool{
	//without regexp
	// arrayDigit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// return slices.Contains(arrayDigit, symbol)

	//with regexp
	regExp := `^\d{1}$`
	re := regexp.MustCompile(regExp)
	return re.MatchString(symbol)

}


func main(){
	symbol := "0"

	fmt.Printf("The symbol %s is a valid digit: %v", symbol, isDigit(symbol))
}