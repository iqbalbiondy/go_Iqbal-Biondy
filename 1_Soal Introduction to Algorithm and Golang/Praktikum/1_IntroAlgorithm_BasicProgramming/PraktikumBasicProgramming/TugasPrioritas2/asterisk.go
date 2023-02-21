package main

import "fmt"

func main() {
	var inputUkuran uint8

	fmt.Println("Sample Test Case")
	fmt.Print("Input: ")
	fmt.Scan(&inputUkuran)

	for i := 0; i < int(inputUkuran); i++ {
		for j := 0; j < int(inputUkuran)-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
