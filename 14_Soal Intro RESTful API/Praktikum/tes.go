package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"1"`
	Id     int    `json:"1"`
	Title  string `json:"IqbalBiondy"`
	Body   string `json:"no body"`
}

func main() {
	// Mengambil semua data postingan
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	// Mengambil data postingan dengan id 3
	resp, err = http.Get("https://jsonplaceholder.typicode.com/posts/3")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	// Membuat data postingan baru
	post := Post{1, 1, "Title", "Body"}
	postBody, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	resp, err = http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
