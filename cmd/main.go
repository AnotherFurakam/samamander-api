package main

import (
	"github.com/AnotherFurakam/samamander-api/internal"
	"github.com/AnotherFurakam/samamander-api/pkg/database"
	"github.com/labstack/echo/v4"
)

func init() {
	database.ConnectToDatabase()
}

func main() {

	e := echo.New()

	internal.Router(e, database.DB)

	e.Logger.Info(e.Start(":5000"))
}
