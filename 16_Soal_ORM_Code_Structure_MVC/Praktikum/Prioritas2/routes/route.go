package routes

import (
	"github.com/iqbalbiondy/go_Iqbal-Biondy/16_Soal_ORM_Code_Structure_MVC/Praktikum/Prioritas2/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// route user
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.POST("/users", controller.CreateUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// route book
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.POST("/books", controller.CreateBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	return e
}
