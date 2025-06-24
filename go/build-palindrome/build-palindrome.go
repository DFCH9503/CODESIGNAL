package main

import(
	"fmt"
	"strings"
)

func splitReverse(s1 string) string{
	runes := []rune(s1)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func buildPalindrome(st string) string{
	i := 0
	aux := []string{}
	for st != splitReverse(st){
		aux = strings.Split(st, "")

		insertionIndex := len(st) - i
		elementToInsert := st[i]
		aux = append(aux[:insertionIndex], append([]string{string(elementToInsert)}, aux[insertionIndex:]...)...)

		st = strings.Join(aux, "")
		i++
	}
	return st
}

func main(){
	st := "abcdc"

	fmt.Printf("The shortest possible string which can be achieved by adding characters to the end of string %s to make it a palindrome is: %s", st, buildPalindrome(st))
}