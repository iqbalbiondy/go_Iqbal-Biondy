package main

import "fmt"

func pow(x, n int) int {
	result := 1
	for n > 0 {
	  if n&1 == 1 {
		result *= x
	  }
	  x *= x
	  n >>= 1
	}
	return result
  }

  func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
 }