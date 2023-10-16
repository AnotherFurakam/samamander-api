package main

import (
	_ "github.com/AnotherFurakam/samamander-api/cmd/docs"
	"github.com/AnotherFurakam/samamander-api/internal"
	"github.com/AnotherFurakam/samamander-api/pkg/database"
	"github.com/AnotherFurakam/samamander-api/pkg/validation"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	database.ConnectToDatabase()
	validation.InitializeValidator()
}

// @title			Samamander API
// @version		1.0
// @description	This is a Samamander API
// @BasePath		/api
func main() {

	e := echo.New()

	internal.Router(e, database.DB)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Info(e.Start(":5000"))
}
