package main

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func commonCharacterCount(s1, s2 string) int{
	res := 0
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, char1 := range s1{
		
		map1[char1] = strings.Count(s1, string(char1))
	}

	for _, char2 := range s2{
		map2[char2] = strings.Count(s2, string(char2))
	}


	for key, value1 := range map1{
		if value2, ok := map2[key]; ok{
			res += min(value1, value2)
		}
	}
	//PROBLEM WITH LARGER STRINGS - STILL-FIXING IT
	// for i := 0 ; i < len(s1) ; i++{
	// 	for j := 0 ; j < len(s2) ; j++{
	// 		if s1[i] == s2[j]{
	// 			res++
	// 			s2 = s2[:j] + s2[j+1:]
	// 			s1 = s1[:i] + s1[i+1:]
	// 		}
	// 		fmt.Printf("s1 - %s, s2 - %s, i - %d, j - %d", s1, s2, i, j)
	// 	}
	// } 
	return res
}

func main(){
	s1 := "aabcc" 
	s2 := "adcaa"
	fmt.Printf("For the strings %s and %s there are %d common characters\n", s1, s2, commonCharacterCount(s1, s2))
}