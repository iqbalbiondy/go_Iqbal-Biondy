package main

import "fmt"

func generate(inputAngka int) [][]int {
    if inputAngka == 0 {
        return [][]int{}
    }
    if inputAngka == 1 {
        return [][]int{{1}}
    }
    barisSebelumnya := generate(inputAngka - 1)
    barisTrakhir := barisSebelumnya[len(barisSebelumnya)-1]
    row := make([]int, inputAngka)
    row[0], row[inputAngka-1] = 1, 1
    for i := 1; i < inputAngka-1; i++ {
        row[i] = barisTrakhir[i-1] + barisTrakhir[i]
    }
    return append(barisSebelumnya, row)
}

func main() {
    inputAngka := 5
    pascal := generate(inputAngka)
    fmt.Println(pascal)
}
