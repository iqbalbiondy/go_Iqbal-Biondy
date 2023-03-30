package main

import (
	"fmt"

)

func ArrayMerge(arrayA, arrayB []string) []string {
    merged := append(arrayA, arrayB...) // menggabungkan array

    mergedLength := len(merged)

    // Menghapus elemen duplikat dengan time complexity logaritmik
    if mergedLength < 2 {
        return merged
    }

    writeIndex := 1
    for readIndex := 1; readIndex < mergedLength; readIndex++ {
        if merged[readIndex] != merged[readIndex-1] {
            merged[writeIndex] = merged[readIndex]
            writeIndex++
        }
    }

    return merged[:writeIndex]
}

func main (){
	fmt.Println(ArrayMerge([]string{"king","devil jin","akuma"},[]string{"eddie","steve","geese"}))
}