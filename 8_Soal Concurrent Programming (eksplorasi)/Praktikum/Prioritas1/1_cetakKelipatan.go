package main

import (
    "fmt"
    "time"
)

func cetakKelipatan(x int) {
    for i := 1; ; i++ {
        if i%x == 0 {
            fmt.Println(i)
        }
        time.Sleep(3 * time.Second)
    }
}

func main() {
    go cetakKelipatan(3)
    for {
        time.Sleep(10 * time.Second)
    }
}
