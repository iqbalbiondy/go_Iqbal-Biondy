package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
                                                              
var users []User

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get user by id",
				"user":    user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if user.Id == id {
			// binding updated data
			updatedUser := new(User)
			if err := c.Bind(updatedUser); err != nil {
				return err
			}
			// update user data
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user",
				"user":    users[i],
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, *user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func main() {
	e := echo.New()

	// routing
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server
	e.Start(":8000")
}
