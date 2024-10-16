package repositories

import (
	"calender-service/config"
	"calender-service/models"
)


type UserRepository struct{}

func (r *UserRepository) Create(user *models.User) error{
	return config.GetDB().Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User,error){
	var user models.User
	err := config.GetDB().First(&user, id).Error
	return &user, err
}
