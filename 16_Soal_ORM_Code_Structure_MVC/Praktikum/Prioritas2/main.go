package main

import (
	"github.com/iqbalbiondy/16_Soal_ORM_Code_Structure_MVC/Praktikum/Prioritas2/config"
	"github.com/iqbalbiondy/16_Soal_ORM_Code_Structure_MVC/Praktikum/Prioritas2/routes" 
)

func main() {
	config.Init()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
