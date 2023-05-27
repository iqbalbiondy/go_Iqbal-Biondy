package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Tampilkan data yang diperoleh dari API")

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

	fmt.Println("Tampilkan data dengan id 3 yang diperoleh dari API")

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

	fmt.Println("Simpan data postingan ke server melalui API")
	
	postData := map[string]string{
		"title":  "iqbalbiondy",
		"body":   "biondy",
		"userId": "1",
	}
	postBody, err := json.Marshal(postData)
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

	fmt.Println("Hapus suatu data melalui API")

	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		panic(err)
	}
	resp, err = http.DefaultClient.Do(req)
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
