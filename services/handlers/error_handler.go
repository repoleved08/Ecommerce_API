package handlers

import "github.com/labstack/echo/v4"

func sendErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, echo.Map{"error": message})
}
