package config

import (
	"github.com/iqbalbiondy/go_Iqbal-Biondy/16_Soal_ORM_Code_Structure_MVC/Praktikum/Prioritas2/models"
	"gorm.io/driver/mysql" 
	"gorm.io/gorm"
)

func Init() {
	initDB()
	initMigrate()
}

var DB *gorm.DB

func initDB() {
	username := "iqbal"
	password := "iqbal"
	host := "127.0.0.1"
	port := "3306"
	dbname := "latihan"
	var err error

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
