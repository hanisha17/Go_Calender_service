package services

import (
	"calender-service/models"

	"github.com/stretchr/testify/mock"
)
type UserServiceInterface interface {
    CreateUser(user *models.User) error
    GetAllUsers() ([]models.User, error)
}


type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) CreateUser(user *models.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserService) GetAllUsers() ([]models.User, error) {
    args := m.Called()
    return args.Get(0).([]models.User), args.Error(1)
}
