package migrator

import (
	mysql_users "Praktikum/repository/mysql/users"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&mysql_users.User{})

}
