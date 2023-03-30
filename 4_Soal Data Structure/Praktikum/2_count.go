package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
    countMap := make(map[string]int)
    for _, str := range slice {
        countMap[str]++
    }
    return countMap
}


func main ()  {
	fmt.Println(Mapping([]string{"asd","qwe","asd","adi","qwe","qwe"}))
	fmt.Println(Mapping([]string{"asd","qwe","asd"}))
	fmt.Println(Mapping([]string{}))
}