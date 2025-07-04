package main

import(
	"fmt"
	"regexp"
	"strconv"
)

func lineEncoding(s string) string{
	re := regexp.MustCompile(`([a-z])\1*`)
	return re.ReplaceAllStringFunc(s, func(it string) string {
		if len(it) > 1 {
			return strconv.Itoa(len(it)) + string(it[0])
		}
		return string(it[0])
	})
}


func main(){
	s := "aabbbc"
	fmt.Printf("The line encondig for the inputString %s is: %s", s, lineEncoding(s))

}