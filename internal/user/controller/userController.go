package controller

import (
	"net/http"
	"strconv"

	"github.com/AnotherFurakam/samamander-api/internal/user/model"
	"github.com/AnotherFurakam/samamander-api/internal/user/service"
	pkg_model "github.com/AnotherFurakam/samamander-api/pkg/pkgModel"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

// GetAll godoc
//
//	@Summary		List users
//	@Description	get users
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Param			pageNumber	query	string	false	"pageNumber"
//	@Param			pageSize	query	string	false	"pageSize"
//	@Router			/user [get]
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
			Message:    "An error occurrerd when querying users: " + err.Error(),
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

// GetById GetAll godoc
//
//	@Summary		Find users
//	@Description	get user by id
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Param			idUser	path	string	false	"idUser"
//	@Router			/user/{idUser} [get]
func (uc *UserController) GetById(c echo.Context) error {
	idUser := c.Param("idUser")

	user, err := uc.userService.FindById(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error occurred while finding user: " + err.Error(),
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
		Message:    "User successfully found",
		Data:       user,
		Success:    true,
		StatusCode: http.StatusOK,
	})
}

// Create User godoc
//
//	@Summary		Create users
//	@Description	create user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			userDto	body	model.CreateUserDto	false	"userDto"
//	@Router			/user [post]
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

// Update User godoc
//
//	@Summary		Update users
//	@Description	Update user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			idUser	path	string				false	"idUser"
//	@Param			userDto	body	model.UpdateUserDto	false	"userDto"
//	@Router			/user/{idUser} [put]
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

// Delete user godoc
//
//	@Summary		Delete user
//	@Description	Delete user by id
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Param			idUser	path	string	false	"idUser"
//	@Router			/user/{idUser} [delete]
func (uc *UserController) Delete(c echo.Context) error {
	idUser := c.Param("idUser")
	user, err := uc.userService.Delete(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_model.ApiResponse[any]{
			Message:    "An error occurred while deleting user: " + err.Error(),
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
