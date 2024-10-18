package repositories

import (
	"calender-service/config"
	"calender-service/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestEventRepository_Create(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	event := &models.Event{
		ID:        1,
		Name:      "Test Event",
		UserID:    1,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(2 * time.Hour),
		RoomID: nil,
	}

	mock.ExpectBegin()
    mock.ExpectExec("INSERT INTO `events`").
        WithArgs(event.UserID, event.Name, event.StartTime, event.EndTime, sqlmock.AnyArg(), event.ID).
        WillReturnResult(sqlmock.NewResult(1, 1)) 
    mock.ExpectCommit()

    repo := EventRepository{}
    err = repo.Create(event)

    assert.NoError(t, err)
    assert.NoError(t, mock.ExpectationsWereMet())
}
