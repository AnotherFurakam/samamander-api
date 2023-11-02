package controller

import (
	"github.com/AnotherFurakam/samamander-api/internal/post/model"
	"github.com/AnotherFurakam/samamander-api/internal/post/service"
	"github.com/AnotherFurakam/samamander-api/pkg/pkgModel"
	"github.com/AnotherFurakam/samamander-api/pkg/utils"
	"github.com/AnotherFurakam/samamander-api/pkg/validation"
	"github.com/labstack/echo/v4"
	"strconv"
)

type PostController struct {
	postService service.PostServiceInterface
}

func NewPostController(postService service.PostServiceInterface) *PostController {
	return &PostController{postService: postService}
}

// GetAll godoc
//
//	@Summary		List Posts
//	@Description	Get Posts
//	@Tags			Post
//	@Accept			*/*
//	@Produce		json
//	@Param			pageNumber	query	string	false	"pageNumber"
//	@Param			pageSize	query	string	false	"pageSize"
//	@Router			/post [get]
func (pc *PostController) GetAll(c echo.Context) error {
	queryPageNumber := c.QueryParam("pageNumber")
	queryPageSize := c.QueryParam("pageSize")

	pageNumber, _ := strconv.Atoi(queryPageNumber)
	pageSize, _ := strconv.Atoi(queryPageSize)

	err := validation.ValidateStruct(pkgModel.PaginationQuery{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	})
	if err != nil {
		return utils.BadRequest(c, "The page number and page size are requierd")
	}

	posts, totalPage, nextPage, prevPage, err := pc.postService.GetAll(pageNumber, pageSize)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	paginationData := pkgModel.PaginationResponse{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		TotalPage:  totalPage,
		NextPage:   nextPage,
		PrevPage:   prevPage,
	}

	return utils.Ok[*[]model.GetPostDto](c, "Posts successfully found", posts, &paginationData)
}

// Create Post godoc
//
//	@Summary		Create post
//	@Description	Create post
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			postDto	body	model.PostDto	false	"postDto"
//	@Router			/post [post]
func (pc *PostController) Create(c echo.Context) error {
	var post model.PostDto
	err := c.Bind(&post)
	if err != nil {
		return utils.BadRequest(c, "Invalid product body")
	}

	postDto, err := pc.postService.Create(&post)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Created[*model.GetPostDto](c, "Post successfully created", postDto)
}
