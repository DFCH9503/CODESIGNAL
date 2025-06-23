package main

import(
	"fmt"
	"regexp"
)

func findEmailDomain(address string) string{

}

func main(){
	address := "prettyandsimple@example.com"

	fmt.Printf("The email domain in the addres %s is : %s", address, findEmailDomain(address))
}