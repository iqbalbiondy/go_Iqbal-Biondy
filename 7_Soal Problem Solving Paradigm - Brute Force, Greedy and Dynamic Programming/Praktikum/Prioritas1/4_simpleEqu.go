package main

import (
    "fmt"
    "math"
    "sort"
)

func SimpleEquations(a, b, c int) {
    var x, y, z float64
    var found bool

    for i := 1; i <= int(math.Sqrt(float64(b))); i++ {
        if b%i == 0 {
            j := b / i
            z = float64(a) - float64(i) - float64(j)
            if z > 0 {
                y = (float64(a)*z - float64(c)) / (2*z - float64(a))
                x = float64(a) - y - z
                if x != y && x != z && y != z {
                    found = true
                    break
                }
            }
        }
    }

    if found {
        results := []int{int(x), int(y), int(z)}
        sort.Ints(results)
        fmt.Println(results)
    } else {
        fmt.Println("no solution")
    }
}

func main() {
    SimpleEquations(1, 2, 3)   // no solution
    SimpleEquations(6, 6, 14)  // [1 2 3]
}
