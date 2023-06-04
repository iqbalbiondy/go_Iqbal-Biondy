package main

import (
	"Praktikum/app/config"
	middlewares "Praktikum/app/middleware"
	"Praktikum/app/routes"
	business_users "Praktikum/businesses/users"
	controller_users "Praktikum/controllers/users"
	migrate "Praktikum/migrator"
	mysql_users "Praktikum/repository/mysql/users"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	migrate.AutoMigrate(db)
	configJWT := middlewares.ConfigJwt{
		SecretJWT: "12345",
	}

	e := echo.New()
	userRepo := mysql_users.NewUsersRepo(db)
	userBusiness := business_users.NewUsersBusiness(userRepo, &configJWT)
	userController := controller_users.NewUsersController(userBusiness)

	routeList := routes.RouteList{
		UserBus: userController,
	}

	routeList.NewRouteList(e)
	e.Logger.Fatal(e.Start(":8080"))
}
