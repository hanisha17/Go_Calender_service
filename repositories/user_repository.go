package repositories

import (
	"calender-service/config"
	"calender-service/models"
)


type UserRepository struct{
	CreateFunc func(user *models.User) error
	GetAllFunc func()([]models.User, error)
	GetAllId func (id uint) (*models.User, error)
}

func (r *UserRepository) Create(user *models.User) error{
	if r.CreateFunc !=nil {
		return r.CreateFunc(user)
	}
	return config.GetDB().Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User,error){
	if r.GetAllId!=nil {
		return r.GetAllId(id)
	}
	var user models.User
	err := config.GetDB().First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetAll() ([]models.User,error){
	if r.GetAllFunc != nil {
		return r.GetAllFunc()
	}
	var users []models.User
	err :=config.GetDB().Find(&users).Error
	return users, err
}
