package routes

import (
	"github.com/JGurus/sin-barreras-backend/api/handlers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, s handlers.UserHandler) {
	handler := handlers.NewUser(s)
	users := e.Group("/users")
	users.POST("", handler.Create)
	users.PUT("/:id", handler.Update)
	users.DELETE("/:id", handler.Delete)
	users.GET("", handler.GetAll)
	users.GET("/:id", handler.GetByID)
}
