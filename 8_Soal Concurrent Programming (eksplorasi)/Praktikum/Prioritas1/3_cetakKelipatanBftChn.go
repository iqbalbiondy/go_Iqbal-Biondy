package main

import (
    "fmt"
    "time"
)

func cetakKelipatan(ch chan<- int) {
    for i := 1; ; i++ {
        if i%3 == 0 {
            ch <- i
        }
		// menunggu selama 3 detik sebelum mengirimkan bilangan berikutnya
        time.Sleep(3 * time.Second)     
	}
}

func main() {
    ch := make(chan int, 5) // buffer channel dengan kapasitas 5
    go cetakKelipatan(ch)
    for {
        fmt.Println(<-ch)
    }
}
