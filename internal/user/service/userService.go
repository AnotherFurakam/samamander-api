package service

import (
	"errors"
	"github.com/AnotherFurakam/samamander-api/internal/user/model"
	"github.com/AnotherFurakam/samamander-api/pkg/utils"
	"github.com/AnotherFurakam/samamander-api/pkg/validation"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserServiceInterface interface {
	GetAll(pageNumber int, pageSize int) (users *[]model.GetUserDto, totalPage *int, nextPage *int, prevPage *int, err error)
	FindById(userId string) (*model.GetUserDto, error)
	Create(userBody model.CreateUserDto) (*model.GetUserDto, error)
	Update(idUser string, userBody *model.UpdateUserDto) (*model.GetUserDto, error)
	Delete(idUser string) (*model.GetUserDto, error)
}

type UserService struct {
	DB *gorm.DB
}

func NewUserService(DB *gorm.DB) *UserService {
	return &UserService{DB: DB}
}

//Private functions

func (u *UserService) findById(idUser string) (*model.User, error) {
	var user *model.User

	result := u.DB.Where("id_user = ?", idUser).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (u *UserService) findByUsername(username string) error {
	var user *model.User
	result := u.DB.Where("user_name", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	if user != nil {
		return errors.New("User with username `" + username + "` already exist")
	}
	return nil
}

func (u *UserService) findByEmail(email string) error {
	var user *model.User
	result := u.DB.Where("email", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	if user != nil {
		return errors.New("User with email `" + email + "` already exist")
	}
	return nil
}

// Service functions

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

func (u *UserService) FindById(idUser string) (*model.GetUserDto, error) {
	err := validation.Validate.Var(idUser, "required,uuid4")
	if err != nil {
		return nil, errors.New("empty or invalid id")
	}

	user, err := u.findById(idUser)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	userDto := model.GetUserDto{
		IdUser:   user.IdUser,
		Username: user.UserName,
		Email:    user.Email,
		IsActive: user.IsActive,
		CreateAt: user.CreateAt,
	}

	return &userDto, nil
}

func (u *UserService) Create(userBody model.CreateUserDto) (*model.GetUserDto, error) {

	err := u.findByUsername(userBody.Username)
	if err != nil {
		return nil, err
	}

	err = u.findByEmail(userBody.Email)
	if err != nil {
		return nil, err
	}

	user := model.User{
		UserName: userBody.Username,
		Email:    userBody.Email,
		Password: userBody.Password,
	}

	err = validation.ValidateStruct(&user)
	if err != nil {
		return nil, err
	}

	result := u.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	userDto := model.GetUserDto{
		IdUser:   user.IdUser,
		Username: user.UserName,
		Email:    user.Email,
		IsActive: user.IsActive,
		CreateAt: user.CreateAt,
	}

	return &userDto, nil
}

func (u *UserService) Update(idUser string, userBody *model.UpdateUserDto) (*model.GetUserDto, error) {
	err := validation.Validate.Var(idUser, "required,uuid4")
	if err != nil {
		return nil, errors.New("empty or invalid id")
	}

	user, err := u.findById(idUser)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if userBody.Username != user.UserName {
		err = u.findByUsername(userBody.Username)
		if err != nil {
			return nil, err
		}
	}

	if userBody.Email != user.Email {
		err = u.findByEmail(userBody.Email)
		if err != nil {
			return nil, err
		}
	}

	user.UserName = userBody.Username
	user.Email = userBody.Email
	user.Password = userBody.Password

	err = validation.ValidateStruct(user)
	if err != nil {
		return nil, err
	}

	result := u.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	userDto := model.GetUserDto{
		IdUser:   user.IdUser,
		Username: user.UserName,
		Email:    user.Email,
		IsActive: user.IsActive,
		CreateAt: user.CreateAt,
	}

	return &userDto, nil

}

func (u *UserService) Delete(idUser string) (*model.GetUserDto, error) {

	err := validation.Validate.Var(idUser, "required,uuid4")
	if err != nil {
		return nil, errors.New("empty or invalid id")
	}

	user, err := u.findById(idUser)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	result := u.DB.Clauses(clause.Returning{}).Where("id_user = ?", idUser).Delete(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	userDto := model.GetUserDto{
		IdUser:   user.IdUser,
		Username: user.UserName,
		Email:    user.Email,
		IsActive: user.IsActive,
		CreateAt: user.CreateAt,
	}
	return &userDto, nil
}
