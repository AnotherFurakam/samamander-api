package controller

import (
	"net/http"
	"strconv"

	"github.com/AnotherFurakam/samamander-api/internal/user/model"
	"github.com/AnotherFurakam/samamander-api/internal/user/service"
	pkg_model "github.com/AnotherFurakam/samamander-api/pkg/model"
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
		return c.JSON(http.StatusBadRequest, pkg_model.ApiResponse[any]{
			Message:    "The page number and page size query params are require",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	//Enter and verify the page number
	pageNumber, err := strconv.Atoi(queryPageNumber)

	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_model.ApiResponse[any]{
			Message:    "Wrong page number, you must enter an integer",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	//Enter ond verify the page number
	pageSize, err := strconv.Atoi(queryPageSize)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_model.ApiResponse[any]{
			Message:    "Wrong page size, you must enter an integer",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	if pageNumber <= 0 || pageSize <= 0 {
		return c.JSON(http.StatusBadRequest, pkg_model.ApiResponse[any]{
			Message:    "Page number and page size cannot be less than 0",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusBadRequest,
		})
	}

	users, totalPage, nextPage, prevPage, err := uc.userService.GetAll(pageNumber, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error ocurrerd when querying users: " + err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, pkg_model.PaginatedApiResponse[*[]model.GetUserDto]{
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

func (uc *UserController) GetById(c echo.Context) error {
	idUser := c.Param("idUser")

	user, err := uc.userService.FindById(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error ocurred while finding user: " + err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, pkg_model.ApiResponse[any]{
			Message:    "User Not Found",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, pkg_model.ApiResponse[*model.GetUserDto]{
		Message:    "User successfully finded",
		Data:       user,
		Success:    true,
		StatusCode: http.StatusOK,
	})
}

func (uc *UserController) Create(c echo.Context) error {
	var user model.CreateUserDto
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "Invalid user body",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}

	userDto, err := uc.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error occurred while creating the user: " + err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusCreated, pkg_model.ApiResponse[*model.GetUserDto]{
		Message:    "User successfully created",
		Data:       userDto,
		Success:    true,
		StatusCode: http.StatusCreated,
	})

}

func (uc *UserController) Update(c echo.Context) error {
	idUser := c.Param("idUser")

	var user model.UpdateUserDto
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "Invalid user body",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}
	userDto, err := uc.userService.Update(idUser, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error occurred while updating the user: " + err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}
	if userDto == nil {
		return c.JSON(http.StatusNotFound, pkg_model.ApiResponse[any]{
			Message:    "User Not Found",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusCreated, pkg_model.ApiResponse[*model.GetUserDto]{
		Message:    "User successfully updated",
		Data:       userDto,
		Success:    true,
		StatusCode: http.StatusOK,
	})
}

func (uc *UserController) Delete(c echo.Context) error {
	idUser := c.Param("idUser")
	user, err := uc.userService.Delete(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error ocurred while deleting user: " + err.Error(),
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusInternalServerError,
		})
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, pkg_model.ApiResponse[any]{
			Message:    "User Not Found",
			Data:       nil,
			Success:    false,
			StatusCode: http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusCreated, pkg_model.ApiResponse[*model.GetUserDto]{
		Message:    "User successfully deleted",
		Data:       user,
		Success:    true,
		StatusCode: http.StatusOK,
	})

}
