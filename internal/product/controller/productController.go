package controller

import "C"
import (
	"github.com/AnotherFurakam/samamander-api/pkg/utils"
	"strconv"

	"github.com/AnotherFurakam/samamander-api/internal/product/model"
	"github.com/AnotherFurakam/samamander-api/internal/product/service"
	"github.com/AnotherFurakam/samamander-api/pkg/pkgModel"
	"github.com/AnotherFurakam/samamander-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.ProductServiceInterface
}

func NewProductController(productService service.ProductServiceInterface) *ProductController {
	return &ProductController{productService: productService}
}

// GetById GetAll godoc
//
//	@Summary		Find products
//	@Description	get product by id
//	@Tags			Product
//	@Accept			*/*
//	@Produce		json
//	@Param			productId	path	string	false	"productId"
//	@Router			/product/{productId} [get]
func (pc *ProductController) GetById(c echo.Context) error {
	productId := c.Param("productId")

	err := validation.Validate.Var(productId, "required")
	if err != nil {
		return utils.BadRequest(c, "Product id is required")
	}

	productDto, err := pc.productService.GetById(productId)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Ok[*model.GetProductDto](c, "Product successfully found", productDto, nil)
}

// GetAll godoc
//
//	@Summary		List products
//	@Description	get products
//	@Tags			Product
//	@Accept			*/*
//	@Produce		json
//	@Param			pageNumber	query	string	false	"pageNumber"
//	@Param			pageSize	query	string	false	"pageSize"
//	@Router			/product [get]
func (pc *ProductController) GetAll(c echo.Context) error {
	queryPageNumber := c.QueryParam("pageNumber")
	queryPageSize := c.QueryParam("pageSize")

	pageNumber, _ := strconv.Atoi(queryPageNumber)
	pageSize, _ := strconv.Atoi(queryPageSize)

	err := validation.Validate.Struct(pkgModel.PaginationQuery{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	})
	if err != nil {
		return utils.BadRequest(c, "The page number and page size are required")
	}

	products, totalPage, nextPage, prevPage, err := pc.productService.GetAll(pageNumber, pageSize)
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
	return utils.Ok[*[]model.GetProductDto](c, "Product successfully found", products, &paginationData)
}

// Create Product godoc
//
//	@Summary		Create product
//	@Description	create product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			productDto	body	model.ProductDto	false	"productDto"
//	@Router			/product [post]
func (pc *ProductController) Create(c echo.Context) error {
	var product model.ProductDto
	err := c.Bind(&product)
	if err != nil {
		return utils.BadRequest(c, "Invalid product body")
	}

	productDto, err := pc.productService.Create(&product)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}

	return utils.Created[*model.GetProductDto](c, "Product successfully created", productDto)
}

// Update Product godoc
//
//	@Summary		Update product
//	@Description	update product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			productId	path	string				false	"productId"
//	@Param			productDto	body	model.ProductDto	false	"productDto"
//	@Router			/product/{productId} [put]
func (pc *ProductController) Update(c echo.Context) error {
	productId := c.Param("productId")
	var product model.ProductDto
	err := c.Bind(&product)
	if err != nil {
		return utils.BadRequest(c, "Invalid product body")
	}

	productDto, err := pc.productService.Update(productId, &product)
	if err != nil {
		return utils.InternalServerError(c, err.Error())
	}
	return utils.Ok[*model.GetProductDto](c, "Product successfully updated", productDto, nil)
}

// Delete Product godoc
//
//	@Summary		Delete product
//	@Description	delete product
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Param			productId	path	string	false	"productId"
//	@Router			/product/{productId} [delete]
func (pc *ProductController) Delete(c echo.Context) error {
	productId := c.Param("productId")

	productDto, err := pc.productService.Delete(productId)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Ok[*model.GetProductDto](c, "Product successfully deleted", productDto, nil)
}
