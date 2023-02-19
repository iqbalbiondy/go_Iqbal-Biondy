package main

import "fmt"

func main() {
	var inputNilai uint8

	fmt.Print("Masukkan nilai anda: ")
	fmt.Scan(&inputNilai)

	for i := 1; i <= int(inputNilai); i++ {
		if inputNilai%3 == 0 && inputNilai%5 == 0 {
			fmt.Println("FizzBuzz")
			break
		} else if inputNilai%3 == 0 {
			fmt.Println("Fizz")
			break
		} else if inputNilai%5 == 0 {
			fmt.Println("Buzz")
			break
		} else {
			print(inputNilai)
			break
		}
	}
}
