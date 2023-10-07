package user

import (
	"github.com/AnotherFurakam/samamander-api/internal/user/controller"
	"github.com/AnotherFurakam/samamander-api/internal/user/model"
	"github.com/AnotherFurakam/samamander-api/internal/user/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserModuleRouter(DB *gorm.DB, group echo.Group) *echo.Group {

	router := group.Group("/users")

	userService := service.NewUserService(DB)
	userController := controller.NewUserController(userService)

	router.GET("", userController.GetAll)

	return router
}

func MigrateModel(DB *gorm.DB) {
	DB.AutoMigrate(&model.User{})
}
