package main

import (
	"fmt"
	"regexp"
)

func longestWord(text string) string {
	pattern := regexp.MustCompile(`[^a-zA-Z]+`)

	words := pattern.Split(text, -1)


	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}

	return longest
}

func main() {
	text := "Ready, steady, go!"
	fmt.Printf("The longest word in the text '%s' is: %s", text, longestWord(text))
}
