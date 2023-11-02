package internal

import (
	"fmt"
	"github.com/AnotherFurakam/samamander-api/internal/post"
	"github.com/AnotherFurakam/samamander-api/internal/product"
	"github.com/AnotherFurakam/samamander-api/internal/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, DB *gorm.DB) {
	v1 := e.Group("api")

	user.ModuleRouter(DB, *v1)
	product.ModuleRouter(DB, *v1)
	post.ModuleRouter(DB, *v1)
}

func Migration(DB *gorm.DB) {
	err := user.MigrateModel(DB)
	if err != nil {
		fmt.Println("An error occurred while migrating the user pkgModel:\n Error: " + err.Error())
	}

	err = product.MigrateModel(DB)
	if err != nil {
		fmt.Println("An error occurred while migrating the product pkgModel:\n Error: " + err.Error())
	}

	err = post.MigrateModel(DB)
	if err != nil {
		fmt.Println("An error occurred while migrating the post pkgModel:\n Error: " + err.Error())
	}

}
