package service

import (
	"github.com/AnotherFurakam/samamander-api/internal/user/model"
	"github.com/AnotherFurakam/samamander-api/pkg/utils"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAll(pageNumber int, pageSize int) (users *[]model.GetUserDto, totalPage *int, nextPage *int, prevPage *int, err error)
	//Create(user *model.CreateUserDto) (*model.GetUserDto, error)
	//Update(user *model.UpdateUserDto) (*model.GetUserDto, error)
	//Delete(userId string) error
}

type UserService struct {
	DB *gorm.DB
}

func NewUserService(DB *gorm.DB) *UserService {
	return &UserService{DB: DB}
}

func (u *UserService) GetAll(pageNumber int, pageSize int) (users *[]model.GetUserDto, totalPage *int, nextPage *int, prevPage *int, err error) {

	var userList []model.User

	offset := (pageNumber - 1) * pageSize

	result := u.DB.Limit(pageSize).Offset(offset).Find(&userList)
	if result.Error != nil {
		return nil, nil, nil, nil, result.Error
	}

	var totalOfRecords int64
	result = u.DB.Model(&userList).Count(&totalOfRecords)
	if result.Error != nil {
		return nil, nil, nil, nil, result.Error
	}

	totalPage, nextPage, prevPage = utils.CalculatePaginationData(pageNumber, pageSize, totalOfRecords)

	var dtoUserList []model.GetUserDto
	for _, user := range userList {
		dtoUserList = append(dtoUserList, model.GetUserDto{
			IdUser:   user.IdUser,
			Username: user.UserName,
			Email:    user.Email,
			IsActive: user.IsActive,
			CreateAt: user.CreateAt,
		})
	}

	return &dtoUserList, totalPage, nextPage, prevPage, result.Error
}
