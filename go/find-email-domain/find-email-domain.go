package main

import(
	"fmt"
	"regexp"
)

func findEmailDomain(address string) string{
	regExpDomain := regexp.MustCompile(`@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,}`)
	domainFound := regExpDomain.FindString(address)
	return domainFound[1:]
}

func main(){
	address := "prettyandsimple@example.com"

	fmt.Printf("The email domain in the addres %s is : %s", address, findEmailDomain(address))
}