package main

import "fmt"

func getMinMax(nilaiVar ...*int) (min, max int) {
    if len(nilaiVar) > 0 {
        min = *nilaiVar[0]
        max = *nilaiVar[0]
    }
    for _, n := range nilaiVar {
        if *n < min {
            min = *n
        }
        if *n > max {
            max = *n
        }
    }
    return min, max
}

func main() {
    var a1, a2, a3, a4, a5, a6, min, max int

    fmt.Scan(&a1)
    fmt.Scan(&a2)
    fmt.Scan(&a3)
    fmt.Scan(&a4)
    fmt.Scan(&a5)
    fmt.Scan(&a6)

    // meneruskan pointer ke variabel untuk fungsi getMinMax
    min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Output")
    fmt.Println("Nilai min:", min)
    fmt.Println("Nilai max:", max)
}