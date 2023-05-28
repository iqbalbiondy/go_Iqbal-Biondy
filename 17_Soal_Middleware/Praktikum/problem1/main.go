package main

import(
	"github.com/iqbalbiondy/17_Soal_Middleware/Praktikum/problem1/config"
	"github.com/iqbalbiondy/17_Soal_Middleware/Praktikum/problem1/routes"
)

func main() {
	config.Init()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}