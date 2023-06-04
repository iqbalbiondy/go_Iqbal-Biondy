package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (DB *gorm.DB) {
	var (
		DBNAME = "latihan_clean"
		DBUSER = "iqbal"
		DBPASS = "iqbalbiondy"
		DBHOST = "iqbaldb.cs8qgr0edqzk.us-east-1.rds.amazonaws.com"
		DBPORT = "3306"
	)

	var err error

	dsn := DBUSER + ":" + DBPASS + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}
