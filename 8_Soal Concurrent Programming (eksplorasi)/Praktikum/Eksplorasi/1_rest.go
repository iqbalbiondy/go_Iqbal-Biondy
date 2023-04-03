package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Produk struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	respon, pesanError := http.Get("https://fakestoreapi.com/products")
	if pesanError != nil {
		panic(pesanError)
	}
	defer respon.Body.Close()

	if respon.StatusCode != http.StatusOK {
		panic(fmt.Errorf("status code tidak valid: %d", respon.StatusCode))
	}

	var produk []Produk
	pesanError = json.NewDecoder(respon.Body).Decode(&produk)
	if pesanError != nil {
		panic(pesanError)
	}

	// membuat channel untuk produk
	produkChannel := make(chan Produk)

	// jalankan fungsi goroutine untuk memproses setiap produk
	for _, product := range produk {
		go func(p Produk) {
			produkChannel <- p
		}(product)
	}

	// mengambil setiap produk dari channel dan menampilkan ke layar
	for i := 0; i < len(produk); i++ {
		fmt.Println("===")
		product := <-produkChannel
		fmt.Println("title:", product.Title)
		fmt.Println("price:", product.Price)
		fmt.Println("category:", product.Category)
		fmt.Println("===")
	}
}
