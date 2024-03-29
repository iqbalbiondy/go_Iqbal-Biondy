package database

import (
	"errors"

	"github.com/iqbalbiondy/17_Soal_Middleware/Praktikum/problem1/config"
	"github.com/iqbalbiondy/17_Soal_Middleware/Praktikum/problem1/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	queryData := config.DB.Find(&books)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBook(id int) ([]models.Book, error) {
	var book []models.Book
	queryData := config.DB.Where("id = ?", id).Find(&book)
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return book, nil
}

func CreateBook(data models.Book) (models.Book, error) {
	queryData := config.DB.Save(&data)
	if e := queryData.Error; e != nil {
		return models.Book{}, e
	}
	return data, nil
}

func UpdateBook(id int, data models.Book) (models.Book, error) {
	queryData := config.DB.Model(&data).Where("id = ?", id).Updates(map[string]interface{}{"id": id, "title": data.Title, "author": data.Author, "publisher": data.Publisher})
	if queryData.RowsAffected == 0 {
		return models.Book{}, errors.New("book not found")
	}
	if e := queryData.Error; e != nil {
		return models.Book{}, e
	}
	return data, nil
}

func DeleteBook(id int) ([]models.Book, error) {
	var book []models.Book
	queryData := config.DB.Unscoped().Where("id = ?", id).Delete(&book)
	if queryData.RowsAffected == 0 {
		return nil, errors.New("book not found")
	}
	if e := queryData.Error; e != nil {
		return nil, e
	}
	return book, nil
}
