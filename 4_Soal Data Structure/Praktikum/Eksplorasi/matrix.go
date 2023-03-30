package main

import (
    "fmt"
    "math"
)

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {9, 8, 9},
    }

    perbedaanDiagonal := hitungBedaDiagonal(matrix)
    fmt.Println("Perbedaan mutlak Diagonal :", perbedaanDiagonal)
}

func hitungBedaDiagonal(matrix [][]int) int {
    n := len(matrix)
    var jumlahDKiri, jumlahDKanan int

    for i := 0; i < n; i++ {
        jumlahDKiri += matrix[i][i]
        jumlahDKanan += matrix[i][n-i-1]
    }

    return int(math.Abs(float64(jumlahDKiri - jumlahDKanan)))
}
