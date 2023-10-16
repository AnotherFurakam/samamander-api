package product

import (
	"github.com/AnotherFurakam/samamander-api/internal/product/controller"
	"github.com/AnotherFurakam/samamander-api/internal/product/model"
	"github.com/AnotherFurakam/samamander-api/internal/product/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ModuleRouter(DB *gorm.DB, group echo.Group) *echo.Group {
	router := group.Group("/product")

	productService := service.NewProductService(DB)
	productController := controller.NewProductController(productService)

	router.GET("", productController.GetAll)
	router.GET("/:productId", productController.GetById)
	router.POST("", productController.Create)
	router.PUT("/:productId", productController.Update)
	router.DELETE("/:productId", productController.Delete)

	return router
}

func MigrateModel(DB *gorm.DB) error {
	err := DB.AutoMigrate(&model.Product{})
	if err != nil {
		return err
	}
	return nil
}
