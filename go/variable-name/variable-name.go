package main

import(
	"fmt"
	"regexp"
)

func variableName(name string)bool{
	matched, _ := regexp.MatchString(`^[a-zA-Z_]+[a-zA-Z0-9_]*$`, name)
	return matched
}

func main(){
	name := "var_1__Int"
	fmt.Printf("the string %s is a valid variable name: %v", name, variableName(name))
}