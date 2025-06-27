package main

import(
	"fmt"
	"regexp"
)

func isMAC48Address(inputString string) bool{
	regExp := `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`
	re := regexp.MustCompile(regExp)
	return re.MatchString(inputString)
}

func main(){
	inputString := "00-1B-63-84-45-E6"

	fmt.Printf("The string %s is a valid standard IEEE 802 MAC-48 address: %v", inputString, isMAC48Address(inputString))
}