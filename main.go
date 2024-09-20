package main

import (
	"github.com/labstack/echo/v4"
	"github.com/repoleved08/ecommerce-api/config"
	//"github.com/repoleved08/ecommerce-api/routes"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/repoleved08/ecommerce-api/docs"

)

func main() {
	// INITIALIZE THE DATABSE FOR GO
	config.InitDB()
	// create new echo instance
	e := echo.New()

	// swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// initialize the routes
	//routes.InitRoutes(e)

	// start server with error handling
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
