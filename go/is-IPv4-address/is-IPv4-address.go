package main

import(
	"fmt"
	"regexp"
)

func isIPv4Address(inputString string) bool{
	pattern := `^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.` +
		`(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.` +
		`(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.` +
		`(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])$`
	res, _ := regexp.MatchString(pattern, inputString)
	return res
}

func main(){
	inputString := "172.16.254.1"

	fmt.Printf("The input %s is an IPv4 valid address: %v", inputString, isIPv4Address(inputString))

}