package main

import(
	"fmt"
)

func stringsRearrangement(inputArray []string) bool {
	permutations := permute(inputArray)
	for _, perm := range permutations {
		if isValidArrangement(perm) {
			return true
		}
	}
	return false
}

func isValidArrangement(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		if diffCount(arr[i], arr[i+1]) != 1 {
			return false
		}
	}
	return true
}

func diffCount(s1, s2 string) int {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}

func permute(arr []string) [][]string {
	var result [][]string
	var generate func(current []string, remaining []string)
	generate = func(current []string, remaining []string) {
		if len(remaining) == 0 {
			result = append(result, current)
			return
		}
		for i, str := range remaining {
			nextRemaining := make([]string, len(remaining)-1)
			copy(nextRemaining[:i], remaining[:i])
			copy(nextRemaining[i:], remaining[i+1:])
			generate(append(current, str), nextRemaining)
		}
	}
	generate([]string{}, arr)
	return result
}


func main(){
	inputArray := []string{"aba", "bbb", "bab"}

	fmt.Printf("It's possible to rearrange the strings, from  %v, in such a way that after the rearrangement the strings at consecutive positions would differ by exactly one character: %v", inputArray, stringsRearrangement(inputArray))
}