package main

import (
    "fmt"
    "time"
)

func printMultiples(ch chan<- int) {
    for i := 1; ; i++ {
        if i%3 == 0 {
            ch <- i
        }
		// menunggu selama 3 detik sebelum mengirimkan bilangan berikutnya
        time.Sleep(3 * time.Second) 
    }
}

func main() {
    ch := make(chan int)
    go printMultiples(ch)
    for {
        fmt.Println(<-ch)
    }
}
