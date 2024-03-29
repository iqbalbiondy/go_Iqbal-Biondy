package database

import (
	"errors"

	"github.com/iqbalbiondy/17_Soal_Middleware/Praktikum/problem1/config"
	"github.com/iqbalbiondy/17_Soal_Middleware/Praktikum/problem1/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUser(id int) ([]models.User, error) {
	var user []models.User
	queryData := config.DB.Where("id = ?", id).Find(&user)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(data models.User) (models.User, error) {
	if e := config.DB.Save(&data).Error; e != nil {
		return models.User{}, e
	}
	return data, nil
}

func UpdateUser(id int, data models.User) (models.User, error) {
	queryData := config.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "name": data.Name, "email": data.Email, "password": data.Password})

	if queryData.RowsAffected == 0 {
		return models.User{}, errors.New("user not found")
	}

	if e := queryData.Error; e != nil {
		return models.User{}, e
	}
	return data, nil
}

func DeleteUser(id int) ([]models.User, error) {
	var user []models.User

	queryData := config.DB.Unscoped().Delete(&models.User{}, id)

	if queryData.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	if e := queryData.Error; e != nil {
		return nil, e
	}
	return user, nil
}
