package main

import "fmt"

func primeX(number int) int {
    hitung := 0
    temp := 2

    for {
        isBilPrima := true
        for i := 2; i < temp; i++ {
            if temp%i == 0 {
                isBilPrima = false
                break
            }
        }

        if isBilPrima {
            hitung++
        }

        if hitung == number {
            return temp
        }

        temp++
    }
}

func main() {
    fmt.Println(primeX(1)) // 2
    fmt.Println(primeX(5)) // 11
    fmt.Println(primeX(8)) // 19
    fmt.Println(primeX(9)) // 23
    fmt.Println(primeX(10)) // 29
}
