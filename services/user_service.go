package services

import (
	"calender-service/models"
	"calender-service/repositories"
	"errors"

	"gorm.io/gorm"
)


type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUserById(id uint) (*models.User,error){
	user , err :=s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user,nil

}