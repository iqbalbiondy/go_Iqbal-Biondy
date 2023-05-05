package main

import (
    "fmt"
    "strings"
)

func main() {
    var input string
    fmt.Print("Masukkan kata: ")
    fmt.Scanln(&input)

    input = strings.ReplaceAll(input, " ", "")

    if isPalindrome(input) {
        fmt.Println("Palindrome")
    } else {
        fmt.Println("Bukan Palindrome")
    }
}

func isPalindrome(s string) bool {
    left := 0
    right := len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
