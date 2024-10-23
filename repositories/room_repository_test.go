package repositories

import (
	"calender-service/config"
	"calender-service/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRoomRepository_Create(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	room := &models.Room{
		ID:   1,
		Name: "Conference Room 1",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `rooms`").
		WithArgs(room.Name, room.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := RoomRepository{}
	err = repo.Create(room)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoomRepository_GetAll(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	expectedRooms := []models.Room{
		{ID: 1, Name: "Conference Room 1"},
		{ID: 2, Name: "Meeting Room 2"},
	}

	mock.ExpectQuery("SELECT \\* FROM `rooms`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(expectedRooms[0].ID, expectedRooms[0].Name).
			AddRow(expectedRooms[1].ID, expectedRooms[1].Name))

	repo := RoomRepository{}
	rooms, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedRooms, rooms)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoomRepository_IsRoomAvailable(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	roomID := uint(1)
	startTime := time.Now()
	endTime := time.Now().Add(2 * time.Hour)

	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `events` WHERE room_id = ?").
		WithArgs(roomID, endTime, startTime, endTime, startTime).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	repo := RoomRepository{}
	isAvailable := repo.IsRoomAvailable(roomID, startTime, endTime)

	assert.True(t, isAvailable)
	assert.NoError(t, mock.ExpectationsWereMet())
}


