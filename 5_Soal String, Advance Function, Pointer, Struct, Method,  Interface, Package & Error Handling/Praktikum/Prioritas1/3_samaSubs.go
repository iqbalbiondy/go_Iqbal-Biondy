package main

import (
    "fmt"
)

func Compare(a, b string) string {
    panjangA := len(a)
    panjangB := len(b)

    // menemukan panjang minimum antara a dan b
    panjangMin := panjangA
    if panjangB < panjangA {
        panjangMin = panjangB
    }

    // melakukan iterasi kedua string dan temukan substring yang cocok
    var substr string
    for i := 0; i < panjangMin; i++ {
        if a[i] == b[i] {
            substr += string(a[i])
        } else {
            break
        }
    }

    return substr
}

func main() {
    fmt.Println(Compare("AKA", "AKASHI")) // Output AKA
    fmt.Println(Compare("KANGOORO", "KANG")) // Output KANG
    fmt.Println(Compare("KI", "KIJANG")) // Output KI
    fmt.Println(Compare("KUPU-KUPU", "KUPU")) // Output KUPU
    fmt.Println(Compare("ILALANG", "ILA")) // Output ILA
}
