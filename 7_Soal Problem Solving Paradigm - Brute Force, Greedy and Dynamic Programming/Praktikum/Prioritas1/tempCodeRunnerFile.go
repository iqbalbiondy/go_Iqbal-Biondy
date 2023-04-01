package main

import (
    "fmt"
    "strconv"
)

func konvertDecKeBinary(n int) string {
    biner := ""
    for n > 0 {
        temp := n % 2
        biner = strconv.Itoa(temp) + biner
        n = n / 2
    }
    return biner
}

func represBiner(n int) []string {
    ans := make([]string, n+1)
    for i := 0; i <= n; i++ {
        ans[i] = konvertDecKeBinary(i)
    }
    return ans
}

func main() {

	m := 2
    ans1 := represBiner(m)
    fmt.Println(ans1)

    n := 3
    ans := represBiner(n)
    fmt.Println(ans)
}
