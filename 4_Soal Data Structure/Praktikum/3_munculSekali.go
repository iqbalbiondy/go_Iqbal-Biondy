package main

import (
    "fmt"
    "strconv"
)

func munculSekali(angka string) []int {
    // Create a map to store the count of each number
    count := make(map[int]int)

    // Loop through the input string and count the occurrences of each number
    for _, char := range angka {
        num, _ := strconv.Atoi(string(char))
        count[num]++
    }

    // Loop through the map and add the numbers that occur only once to the result slice
    result := []int{}
    for num, c := range count {
        if c == 1 {
            result = append(result, num)
        }
    }

    return result
}

func main()  {
	fmt.Println(munculSekali("1234123"))	
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}