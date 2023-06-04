package routes

import (
	controller_users "Praktikum/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteList struct {
	UserBus controller_users.ControllerUser
}

func (rl *RouteList) NewRouteList(e *echo.Echo) {
	e.GET("/users", rl.UserBus.GetAllUser)
	e.POST("/users", rl.UserBus.Register)
	e.POST("/users/login", rl.UserBus.Login)
}
