package service

import (
	"github.com/AnotherFurakam/samamander-api/internal/post/model"
	"github.com/AnotherFurakam/samamander-api/pkg/utils"
	"github.com/AnotherFurakam/samamander-api/pkg/validation"
	"gorm.io/gorm"
)

type PostServiceInterface interface {
	Create(postBody *model.PostDto) (postDto *model.GetPostDto, err error)
	GetAll(pageNumber int, pageSize int) (posts *[]model.GetPostDto, totalPage *int, nextPage *int, prevPage *int, err error)
}

type PostService struct {
	DB *gorm.DB
}

func NewPostService(DB *gorm.DB) *PostService {
	return &PostService{DB: DB}
}

// Service functions

func (ps *PostService) Create(postBody *model.PostDto) (postDto *model.GetPostDto, err error) {
	err = validation.ValidateStruct(postBody)
	if err != nil {
		return nil, err
	}

	post := model.Post{
		Title: postBody.Title,
		Body:  postBody.Body,
	}

	result := ps.DB.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	postDto = &model.GetPostDto{
		IdPost:   post.IdPost,
		Title:    post.Title,
		Body:     post.Body,
		CreateAt: post.CreateAt,
		IsActive: post.IsActive,
	}

	return postDto, nil
}

func (ps *PostService) GetAll(pageNumber int, pageSize int) (posts *[]model.GetPostDto, totalPage *int, nextPage *int, prevPage *int, err error) {
	var postsList []model.Post

	offset := (pageNumber - 1) * pageSize

	result := ps.DB.Limit(pageSize).Offset(offset).Find(&postsList)
	if result.Error != nil {
		return nil, nil, nil, nil, result.Error
	}

	var totalOfRecords int64
	result = ps.DB.Model(&postsList).Count(&totalOfRecords)
	if result.Error != nil {
		return nil, nil, nil, nil, result.Error
	}

	totalPage, nextPage, prevPage = utils.CalculatePaginationData(pageNumber, pageSize, totalOfRecords)

	var dtoPostList []model.GetPostDto
	for _, post := range postsList {
		dtoPostList = append(dtoPostList, model.GetPostDto{
			IdPost:   post.IdPost,
			Title:    post.Title,
			Body:     post.Body,
			IsActive: post.IsActive,
			CreateAt: post.CreateAt,
		})
	}
	return &dtoPostList, totalPage, nextPage, prevPage, nil
}
