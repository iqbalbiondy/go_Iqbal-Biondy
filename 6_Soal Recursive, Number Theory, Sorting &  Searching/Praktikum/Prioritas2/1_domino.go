package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	// Membuat array untuk menyimpan kemungkinan kartu yang bisa dimainkan
	var peluang [][]int

	// Memeriksa setiap kartu di tangan
	for _, card := range cards {
		// Jika salah satu sisi kartu sama dengan kartu di deck, kartu tersebut bisa dimainkan
		if card[0] == deck[0] || card[1] == deck[0] || card[0] == deck[1] || card[1] == deck[1] {
			// Tambahkan kartu tersebut ke dalam array kemungkinan kartu yang bisa dimainkan
			peluang = append(peluang, card)
		}
	}

	// Jika tidak ada kartu yang bisa dimainkan, keluarkan pesan 'tutup kartu'
	if len(peluang) == 0 {
		return "tutup kartu"
	}

	// Mencari kartu yang jumlah titiknya terbanyak
	maxPoints := 0
	var maxCard []int
	for _, card := range peluang {
		points := card[0] + card[1]
		if points > maxPoints {
			maxPoints = points
			maxCard = card
		}
	}

	// Mengembalikan kartu yang jumlah titiknya terbanyak
	return maxCard
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1})) // tutup kartu
}
