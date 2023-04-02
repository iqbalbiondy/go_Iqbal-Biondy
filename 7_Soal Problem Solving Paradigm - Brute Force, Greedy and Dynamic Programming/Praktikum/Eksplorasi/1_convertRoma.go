package main

import "fmt"

func konverIntRoma(inputanNilai int) string {
    nilaiRoma := []struct {
        value  int
        simbol string
    }{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    result := ""
    for _, r := range nilaiRoma {
        for inputanNilai >= r.value {
            result += r.simbol
            inputanNilai -= r.value
        }
    }

    return result
}

func main() {
    fmt.Println(konverIntRoma(4))     // Output: IV
    fmt.Println(konverIntRoma(9))     // Output: IX
    fmt.Println(konverIntRoma(23))    // Output: XXIII
    fmt.Println(konverIntRoma(2021))  // Output: MMXXI
    fmt.Println(konverIntRoma(1646))  // Output: MDCXLVI
}
