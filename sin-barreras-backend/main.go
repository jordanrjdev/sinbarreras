package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JGurus/sin-barreras-backend/api/routes"
	"github.com/JGurus/sin-barreras-backend/config"
	"github.com/JGurus/sin-barreras-backend/database"
	"github.com/JGurus/sin-barreras-backend/database/implementation"
	"github.com/JGurus/sin-barreras-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configServer, err := config.GetConfigServer()
	if err != nil {
		log.Fatalf(err.Error())
	}
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	service := service.NewUser(implementation.NewUserImpl(database.GetConnection()))
	routes.UserRoutes(e, service)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	//e.Static("/avatar", "assets/images/avatar")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", configServer.Port)))
}
