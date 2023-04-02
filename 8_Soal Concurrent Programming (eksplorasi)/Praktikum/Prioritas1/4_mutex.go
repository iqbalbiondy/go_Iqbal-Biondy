package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var angkaGlobalVar []int
var mutex sync.Mutex

func tambahAngkaRandom(wg *sync.WaitGroup) {
    defer wg.Done()

    mutex.Lock()
    defer mutex.Unlock()

    rand.Seed(time.Now().UnixNano())
    angkaLocalVar := rand.Intn(100)
    angkaGlobalVar = append(angkaGlobalVar, angkaLocalVar)

    fmt.Printf("Menambahkan %d ke dalam slice\n", angkaLocalVar)
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go tambahAngkaRandom(&wg)
    }

    wg.Wait()

    fmt.Printf("Slice Final yaitu : %v\n", angkaGlobalVar)
}
