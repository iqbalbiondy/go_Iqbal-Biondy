package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	var1 := len(jumps)
	temp := make([]int, var1)
	temp[0] = 0
	temp[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < var1; i++ {
		temp[i] = int(math.Min(float64(temp[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))), float64(temp[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))))
	}

	return temp[var1-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // output 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // output 40
}
