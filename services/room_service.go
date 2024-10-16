package services

import (
	"calender-service/models"
	"calender-service/repositories"
)


type RoomService struct {
	repo *repositories.RoomRepository
}

func NewRoomService(repo *repositories.RoomRepository) *RoomService {
	return &RoomService{repo}
}

func (s *RoomService) CreateRoom(room *models.Room) error {
	return s.repo.Create(room)
}

func (s *RoomService) GetAll() ([]models.Room,error){
	return s.repo.GetAll()
}