package post

import (
	"github.com/AnotherFurakam/samamander-api/internal/post/controller"
	"github.com/AnotherFurakam/samamander-api/internal/post/model"
	"github.com/AnotherFurakam/samamander-api/internal/post/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ModuleRouter(DB *gorm.DB, group echo.Group) *echo.Group {
	router := group.Group("/post")

	postService := service.NewPostService(DB)
	postController := controller.NewPostController(postService)

	router.GET("", postController.GetAll)
	router.POST("", postController.Create)
	router.PUT("/:postId", postController.Update)
	router.DELETE("/:postId", postController.Delete)

	return router

}

func MigrateModel(DB *gorm.DB) error {
	err := DB.AutoMigrate(&model.Post{})
	if err != nil {
		return err
	}
	return nil
}
