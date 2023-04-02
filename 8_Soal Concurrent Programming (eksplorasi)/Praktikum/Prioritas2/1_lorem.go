package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hitungHurufParalel(huruf string, integer int) map[rune]int {
	// Membuat map untuk menyimpan frekuensi dari setiap huruf
	frekuensi := make(map[rune]int)

	// Membuat channel untuk menerima chunk dari teks
	chunks := make(chan string, integer)

	// Membagi teks menjadi beberapa chunk dan mengirimkan ke channel
	go func() {
		defer close(chunks)
		chunkSize := (len(huruf) + integer - 1) / integer
		for i := 0; i < len(huruf); i += chunkSize {
			end := i + chunkSize
			if end > len(huruf) {
				end = len(huruf)
			}
			chunks <- huruf[i:end]
		}
	}()

	// Membuat channel untuk menerima hasil perhitungan dari setiap worker
	hasilHitung := make(chan map[rune]int, integer)

	// Membuat sejumlah worker untuk menghitung frekuensi pada setiap chunk
	var wg sync.WaitGroup
	wg.Add(integer)
	for i := 0; i < integer; i++ {
		go func() {
			defer wg.Done()
			// Membuat map lokal untuk menyimpan frekuensi dari setiap huruf pada chunk
			localFreq := make(map[rune]int)
			for chunk := range chunks {
				for _, char := range chunk {
					localFreq[char]++
				}
			}
			// Mengirimkan hasil perhitungan frekuensi dari worker ke channel hasilHitung
			hasilHitung <- localFreq
		}()
	}

	// Mengumpulkan hasil perhitungan dari semua worker
	go func() {
		wg.Wait()
		close(hasilHitung)
	}()

	// Menggabungkan hasil perhitungan dari setiap worker menjadi satu map
	for freq := range hasilHitung {
		for char, count := range freq {
			frekuensi[char] += count
		}
	}

	return frekuensi
}

func main() {
	// Menentukan jumlah worker
	integer := runtime.NumCPU()

	// Menghitung frekuensi pada teks
	huruf := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	frekuensi := hitungHurufParalel(huruf, integer)

	// Menampilkan hasil perhitungan frekuensi pada setiap huruf
	for char, count := range frekuensi {
		fmt.Printf("%c:%d\n", char, count)
	}
}
