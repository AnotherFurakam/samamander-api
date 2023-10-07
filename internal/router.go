package internal

import (
	"github.com/AnotherFurakam/samamander-api/internal/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, DB *gorm.DB) {
	v1 := e.Group("api")

	user.UserModuleRouter(DB, *v1)
}

func Migration(DB *gorm.DB) {
	user.MigrateModel(DB)
}
