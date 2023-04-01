package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	valueMax, valueFar := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		valueMax = max(arr[i], valueMax+arr[i])
		valueFar = max(valueFar, valueMax)
	}
	return valueFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
  	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}
