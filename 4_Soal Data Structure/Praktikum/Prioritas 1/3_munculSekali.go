package main

import (
    "fmt"
    "strconv"
)

func munculSekali(angka string) []int {
 // Buat map untuk menyimpan hitungan setiap angka
    count := make(map[int]int)

    // perulangan string input dan hitung kemunculan setiap angka   
    for _, char := range angka {
        num, _ := strconv.Atoi(string(char))
        count[num]++
    }
// perulangan map dan tambahkan angka yang muncul hanya sekali ke potongan hasil
    result := []int{}
    for num, c := range count {
        if c == 1 {
            result = append(result, num)
        }
    }

    return result
}

func main()  {
	fmt.Println(munculSekali("1234123"))	
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}