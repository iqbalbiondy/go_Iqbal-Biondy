package main

import (
	"fmt"
	"sort"
)

type pair struct {
	str  string
	count int
}

func MostAppearItem(itemInput []string) []pair {
	hitung := make(map[string]int)
	for _, item := range itemInput {
		hitung[item]++
	}

	pairs := make([]pair, len(hitung))
	i := 0
	for str, count := range hitung {
		pairs[i] = pair{str, count}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // [{golang 1} {ruby 2} {js 4}]
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // [{C 1} {D 2} {B 3} {A 4}]
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"})) // [{football 1} {basketball 1} {tenis 1}]
}
