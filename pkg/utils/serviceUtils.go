package utils

import (
	"github.com/AnotherFurakam/samamander-api/pkg/pkgModel"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func FindModelByField(DB *gorm.DB, model any, field string, value string) (err error) {
	result := DB.Where(field, value).First(model)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func BadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, pkgModel.ApiResponse[any]{
		Message:    "An error occurred while creating: " + message,
		Data:       nil,
		Success:    false,
		StatusCode: http.StatusBadRequest,
	})
}

func InternalServerError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, pkgModel.ApiResponse[any]{
		Message:    "An error occurred while searching: " + message,
		Data:       nil,
		Success:    false,
		StatusCode: http.StatusInternalServerError,
	})
}

func Ok[T any](c echo.Context, message string, data T, paginationData *pkgModel.PaginationResponse) error {
	if paginationData == nil {
		return c.JSON(http.StatusOK, pkgModel.ApiResponse[T]{
			Message:    message,
			Data:       data,
			Success:    true,
			StatusCode: http.StatusOK,
		})
	}
	return c.JSON(http.StatusOK, pkgModel.PaginatedApiResponse[T]{
		Message:    "List of products successfully obtained",
		Data:       data,
		Success:    true,
		StatusCode: http.StatusOK,
		PageNumber: paginationData.PageNumber,
		PageSize:   paginationData.PageSize,
		TotalPage:  paginationData.TotalPage,
		NextPage:   paginationData.NextPage,
		PrevPage:   paginationData.PrevPage,
	})
}

func Created[T any](c echo.Context, message string, data T) error {
	return c.JSON(http.StatusCreated, pkgModel.ApiResponse[T]{
		Message:    message,
		Data:       data,
		Success:    true,
		StatusCode: http.StatusCreated,
	})
}
