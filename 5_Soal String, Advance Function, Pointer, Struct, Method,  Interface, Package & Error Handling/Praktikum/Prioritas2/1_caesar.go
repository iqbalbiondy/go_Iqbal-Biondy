package main

import "fmt"

func caesar(offset int, input string) string {
	shift := rune(offset % 26)
	runes := []rune(input)
	for i, r := range runes {
		if r != ' ' {
			// Hitung nilai rune yang digeser
			geser := r + shift
			// Wrap untuk nilai di luar alfabet huruf kecil
			if geser > 'z' {
				geser -= 26
			}
			runes[i] = geser
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
