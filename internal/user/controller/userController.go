package controller

import (
	"net/http"
	"strconv"

	"github.com/AnotherFurakam/samamander-api/internal/user/model"
	"github.com/AnotherFurakam/samamander-api/internal/user/service"
	"github.com/AnotherFurakam/samamander-api/pkg/models"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetAll(c echo.Context) error {
	queryPageNumber := c.QueryParam("pageNumber")
	queryPageSize := c.QueryParam("pageSize")

	if len(queryPageNumber) <= 0 || len(queryPageSize) <= 0 {
		return c.JSON(http.StatusBadRequest, models.ApiResponse[any]{
			Message:    "The page number and page size query params are require",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	//Enter and verify the page number
	pageNumber, err := strconv.Atoi(queryPageNumber)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ApiResponse[any]{
			Message:    "Wrong page number, you must enter an integer",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	//Enter ond verify the page number
	pageSize, err := strconv.Atoi(queryPageSize)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ApiResponse[any]{
			Message:    "Wrong page size, you must enter an integer",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	if pageNumber <= 0 || pageSize <= 0 {
		return c.JSON(http.StatusBadRequest, models.ApiResponse[any]{
			Message:    "Page number and page size cannot be less than 0",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	users, totalPage, nextPage, prevPage, err := uc.userService.GetAll(pageNumber, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ApiResponse[any]{
			Message:    "An error ocurrerd when queryng the data: " + err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.PaginatedApiResponse[*[]model.GetUserDto]{
		Message:    "List of users successfully obtained",
		Data:       users,
		Success:    true,
		StatusCode: http.StatusOK,
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalPage:  totalPage,
		NextPage:   nextPage,
		PrevPage:   prevPage,
	})
}
