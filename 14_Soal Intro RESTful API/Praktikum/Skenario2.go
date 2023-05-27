package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0"
)

type Book struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Publisher string `json:"publisher"`
}

func main() {
	// GET request
	resp, err := http.Get(baseURL + "/books")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("GET Response Status:", resp.Status)
	books := make([]Book, 0)
	err = json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Books: %+v\n\n", books)

	// POST request
	book := Book{
		Title:    "Golang Basics",
		Author:   "John Doe",
		Publisher: "ABC Publishing",
	}
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp, err = http.Post(baseURL + "/books", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("POST Response Status:", resp.Status)
	var newBook Book
	err = json.NewDecoder(resp.Body).Decode(&newBook)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("New Book: %+v\n\n", newBook)

	// PUT request
	updatedBook := Book{
		Title:    "Golang Intermediate",
		Author:   "John Doe",
		Publisher: "ABC Publishing",
	}
	jsonData, err = json.Marshal(updatedBook)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req, err := http.NewRequest(http.MethodPut, baseURL+"/books/1", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("PUT Response Status:", resp.Status)
	var updated bool
	err = json.NewDecoder(resp.Body).Decode(&updated)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Updated: %+v\n\n", updated)

	// DELETE request
	req, err = http.NewRequest(http.MethodDelete, baseURL+"/books/1", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("DELETE Response Status:", resp.Status)
	var deleted bool
	err = json.NewDecoder(resp.Body).Decode(&deleted)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Deleted: %+v\n\n", deleted)
}
