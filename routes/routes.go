package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/ecommerce-api/services/handlers"
)


func InitRoutes(e *echo.Echo) {
	e.POST("/api/products/create", handlers.ProductHandler)
}
