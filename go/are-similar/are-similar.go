package main

import(
	"fmt"
)

func areSimilar(a, b []int)bool{
	sw := 0
	var ar int
	var br int
	ret := true
	for idx, val := range a{
		if val == b[idx]{
			continue
		}
		if sw == 0{
			sw++
			ar = val
			br = b[idx]
			ret = false
		}else{
			if ar != b[idx] || br != val{
				sw++
			}else{
				ret = true
				sw++
			}
		}

	}
	return ret && (sw == 0 || sw == 2)
}

func main(){
	a := []int{1, 2, 2} 
	b := []int{2, 1, 1} 
	fmt.Printf("The answer for the arrays %v and %v is %v", a, b, areSimilar(a, b))
}