package main

import "fmt"

func main() {
	var bilangan int8

	fmt.Print("Masukkan bilangan yang akan dicek : ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 {
		fmt.Println("Anda menginputkan ", bilangan, " dan termasuk bilangan genap")
	} else {
		fmt.Println("Anda menginputkan ", bilangan, " dan termasuk bilangan ganjil")
	}
}
