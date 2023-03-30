package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
    // membuat Map yang dapat menyimpan nilai unik
    unique := make(map[string]bool)

    // menambahkan semua value array A ke map
    for _, value := range arrayA {
        unique[value] = true
    }

    // menambahkan semua value array B ke map jika nilai tsb blm ada
    for _, value := range arrayB {
        if _, ok := unique[value]; !ok {
            unique[value] = true
        }
    }

    // membuat slice baru untuk value unik
    merged := make([]string, 0, len(unique))
    for key := range unique {
        merged = append(merged, key)
    }

    return merged
}


func main (){
	fmt.Println(ArrayMerge([]string{"king","devil jin","akuma"},[]string{"eddie","steve","geese"}))
    fmt.Println(ArrayMerge([]string{"sergei","jin"},[]string{"jin","steve","bryan"}))
    fmt.Println(ArrayMerge([]string{"alisa","yoshimitsu"},[]string{"devil jin","yoshimitsu","alisa","law"}))
    fmt.Println(ArrayMerge([]string{},[]string{"devil jin","sergei"}))
    fmt.Println(ArrayMerge([]string{"hwoarang"},[]string{}))
    fmt.Println(ArrayMerge([]string{},[]string{}))
}