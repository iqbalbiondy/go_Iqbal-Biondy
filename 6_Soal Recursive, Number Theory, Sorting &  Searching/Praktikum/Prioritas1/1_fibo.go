package main

import "fmt"

func fibonacci(inputanNilai int) int {
    if inputanNilai == 0 {
        return 0
    }

    temp1 := 0
    temp2 := 1
    result := 1

    for i := 2; i <= inputanNilai; i++ {
        result = temp1 + temp2
        temp1 = temp2
        temp2 = result
    }

    return result
}

func main() {
    fmt.Println(fibonacci(0))  // 0
    fmt.Println(fibonacci(2))  // 1
    fmt.Println(fibonacci(9))  // 34
    fmt.Println(fibonacci(10)) // 55
    fmt.Println(fibonacci(12)) // 144
}
