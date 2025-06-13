package main

import(
	"fmt"
	"math"
)

func max(a, b int) int{
	if a < b {
		return b
	}
	return a
}
func growingPlant(upSpeed, downSpeed, desiredHeight int) int{
	grownSpeed := upSpeed - downSpeed
	return max(0, int(math.Ceil(float64((desiredHeight - upSpeed) / grownSpeed)))) + 1
}

func main(){
	upSpeed, downSpeed, desiredHeight := 100, 10, 910

	fmt.Printf("The plant will reach the desired height %d in %d day/s", desiredHeight,growingPlant(upSpeed, downSpeed, desiredHeight))
}

