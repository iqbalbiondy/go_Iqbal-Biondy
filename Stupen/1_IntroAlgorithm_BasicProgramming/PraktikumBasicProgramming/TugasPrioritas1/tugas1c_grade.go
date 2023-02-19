package main

import "fmt"



func main() {
	var nilai float32
	
	fmt.Print("Masukkan nilai anda : ")
	fmt.Scan(&nilai)

	if nilai<0||nilai>100 {
		fmt.Println("Nilai Invalid")
	}

	if nilai>=80&&nilai<=100 {
		fmt.Println("Grade A")
	}else if nilai>=65&&nilai<80 {
		fmt.Println("Grade B")
	}else if nilai>=50&&nilai<65{
		fmt.Println("Grade C")
	}else if nilai>35&&nilai<50 {
		fmt.Println("Grade D")
	}else if nilai>=0 && nilai<35 {
		fmt.Println("Grade E")
	}
}