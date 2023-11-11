package service

import (
	"errors"
	"github.com/AnotherFurakam/samamander-api/internal/product/model"
	"github.com/AnotherFurakam/samamander-api/pkg/utils"
	"github.com/AnotherFurakam/samamander-api/pkg/validation"
	"gorm.io/gorm"
)

type ProductServiceInterface interface {
	GetById(productId string) (productDto *model.GetProductDto, err error)
	GetAll(pageNumber int, pageSize int) (products *[]model.GetProductDto, totalPage *int, nextPage *int, prevPage *int, err error)
	Create(productBody *model.ProductDto) (productDto *model.GetProductDto, err error)
	Update(productId string, productBody *model.ProductDto) (productDto *model.GetProductDto, err error)
	Delete(productId string) (productDto *model.GetProductDto, err error)
}

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(DB *gorm.DB) *ProductService {
	return &ProductService{DB: DB}
}

// Private functions

//Service functions

func (ps *ProductService) GetById(productId string) (productDto *model.GetProductDto, err error) {
	var product *model.Product
	var errorMessage string

	err = utils.FindModelByField(ps.DB, &product, "id_product", productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage = "Product not found"
			return nil, errors.New(errorMessage)
		}
		return nil, err
	}

	productDto = &model.GetProductDto{
		IdProduct:   product.IdProduct,
		Name:        product.Name,
		Description: product.Description,
		UrlImage:    product.UrlImage,
		IsActive:    product.IsActive,
		CreateAt:    product.CreateAt,
	}

	return productDto, nil
}

func (ps *ProductService) GetAll(pageNumber int, pageSize int) (products *[]model.GetProductDto, totalPage *int, nextPage *int, prevPage *int, err error) {

	var productList []model.Product

	offset := (pageNumber - 1) * pageSize

	result := ps.DB.Limit(pageSize).Offset(offset).Find(&productList)
	if result.Error != nil {
		return nil, nil, nil, nil, result.Error
	}

	var totalOfRecords int64
	result = ps.DB.Model(&productList).Count(&totalOfRecords)
	if result.Error != nil {
		return nil, nil, nil, nil, result.Error
	}

	totalPage, nextPage, prevPage = utils.CalculatePaginationData(pageNumber, pageSize, totalOfRecords)

	var dtoUserList []model.GetProductDto
	for _, product := range productList {
		dtoUserList = append(dtoUserList, model.GetProductDto{
			IdProduct:   product.IdProduct,
			Name:        product.Name,
			Description: product.Description,
			UrlImage:    product.UrlImage,
			IsActive:    product.IsActive,
			CreateAt:    product.CreateAt,
		})
	}

	return &dtoUserList, totalPage, nextPage, prevPage, nil
}

func (ps *ProductService) Create(productBody *model.ProductDto) (productDto *model.GetProductDto, err error) {
	err = validation.ValidateStruct(productBody)
	if err != nil {
		return nil, err
	}

	var productFound *model.Product
	err = utils.FindModelByField(ps.DB, &productFound, "name", productBody.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if productFound.Name == productBody.Name {
		errorMessage := "The product with name " + productBody.Name + " already exist."
		return nil, errors.New(errorMessage)
	}

	product := model.Product{
		Name:        productBody.Name,
		UrlImage:    productBody.UrlImage,
		Description: productBody.Description,
	}

	result := ps.DB.Save(&product)
	if result.Error != nil {
		return nil, err
	}

	productDto = &model.GetProductDto{
		IdProduct:   product.IdProduct,
		Name:        product.Name,
		Description: product.Description,
		UrlImage:    product.UrlImage,
		IsActive:    product.IsActive,
		CreateAt:    product.CreateAt,
	}

	return productDto, nil
}

func (ps *ProductService) Update(productId string, productBody *model.ProductDto) (productDto *model.GetProductDto, err error) {
	err = validation.Validate.Struct(productBody)
	if err != nil {
		return nil, err
	}

	var productFound *model.Product
	var errorMessage string
	err = utils.FindModelByField(ps.DB, &productFound, "id_product", productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage = "Product not found"
			return nil, errors.New(errorMessage)
		}
		return nil, err
	}

	// Think about login for this section
	err = utils.FindModelByField(ps.DB, productFound, "name", productBody.Name)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) == false {
		return nil, err
	}
	if len(productFound.Name) > 0 && productFound.Name == productBody.Name && productFound.IdProduct.String() != productId {
		errorMessage = "The product with name " + productBody.Name + " already exist."
		return nil, errors.New(errorMessage)
	}

	product := &model.Product{
		Name:        productBody.Name,
		UrlImage:    productBody.UrlImage,
		Description: productBody.Description,
	}
	result := ps.DB.Save(product)
	if result.Error != nil {
		return nil, err
	}
	productDto = &model.GetProductDto{
		IdProduct:   product.IdProduct,
		Name:        product.Name,
		Description: product.Description,
		UrlImage:    product.UrlImage,
		IsActive:    product.IsActive,
		CreateAt:    product.CreateAt,
	}
	return productDto, nil
}

func (ps *ProductService) Delete(productId string) (productDto *model.GetProductDto, err error) {
	var errorMessage string
	err = validation.Validate.Var(productId, "required")
	if err != nil {
		errorMessage = "The product id is required"
		return nil, errors.New(errorMessage)
	}

	var productById model.Product
	err = utils.FindModelByField(ps.DB, &productById, "id_product", productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage = "Product not found"
			return nil, errors.New(errorMessage)
		}
		return nil, err
	}

	result := ps.DB.Delete(&productById)
	if result.Error != nil {
		return nil, result.Error
	}

	productDto = &model.GetProductDto{
		IdProduct:   productById.IdProduct,
		Name:        productById.Name,
		Description: productById.Description,
		UrlImage:    productById.UrlImage,
		IsActive:    productById.IsActive,
		CreateAt:    productById.CreateAt,
	}

	return productDto, nil
}
