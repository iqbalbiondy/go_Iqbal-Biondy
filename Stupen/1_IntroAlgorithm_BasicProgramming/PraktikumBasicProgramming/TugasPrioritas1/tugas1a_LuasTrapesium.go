package main

import "fmt"

func main() {
	var a, b, t, luas float64

	fmt.Println("Masukkan nilai a : ")
	fmt.Scan(&a)

	fmt.Println("Masukkan nilai b : ")
	fmt.Scan(&b)

	fmt.Println("Masukkan nilai t : ")
	fmt.Scan(&t)

	luas = 0.5 * (a + b) * t
	fmt.Println("Luas Trapesium anda adalah ", luas)

}
