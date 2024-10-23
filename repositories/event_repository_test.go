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

func TestEventRepository_Update(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	event := &models.Event{
		ID:        1,
		Name:      "Updated Event",
		UserID:    1,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(2 * time.Hour),
		RoomID:    nil,
	}

	mock.ExpectBegin()

	mock.ExpectExec("UPDATE `events`").
		WithArgs(event.UserID, event.Name, event.StartTime, event.EndTime, sqlmock.AnyArg(), event.ID).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate successful update

	mock.ExpectCommit()

	repo := EventRepository{}

	err = repo.Update(event)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

// func TestEventRepository_GetByUserAndDateRange(t *testing.T) {
// 	// Set up the mock DB and sqlmock
// 	db, mock, err := config.SetupMockDB()
// 	assert.NoError(t, err)
// 	config.SetDB(db)

// 	// Define the inputs and expected output
// 	userID := uint(1)
// 	start := time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC)
// 	end := time.Date(2024, 10, 21, 0, 0, 0, 0, time.UTC)

// 	// Expected events to be returned
// 	expectedEvents := []models.Event{
// 		{
// 			ID:        1,
// 			Name:      "Event 1",
// 			UserID:    userID,
// 			StartTime: time.Now(),
// 			EndTime:   time.Now().Add(2 * time.Hour),
// 		},
// 		{
// 			ID:        2,
// 			Name:      "Event 2",
// 			UserID:    userID,
// 			StartTime: time.Now().Add(3 * time.Hour),
// 			EndTime:   time.Now().Add(4 * time.Hour),
// 		},
// 	}

// 	// Mock the SQL SELECT query
// 	mock.ExpectQuery("SELECT \\* FROM `events` WHERE user_id = ? AND").
// 		WithArgs(userID, start, end, start, end).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "user_id", "start_time", "end_time"}).
// 			AddRow(expectedEvents[0].ID, expectedEvents[0].Name, expectedEvents[0].UserID, expectedEvents[0].StartTime, expectedEvents[0].EndTime).
// 			AddRow(expectedEvents[1].ID, expectedEvents[1].Name, expectedEvents[1].UserID, expectedEvents[1].StartTime, expectedEvents[1].EndTime))

// 	// Create the repository instance
// 	repo := EventRepository{}

// 	// Call the GetByUserAndDateRange method
// 	events, err := repo.GetByUserAnddateRange(userID, start, end)
	
// 	// Assert that no errors occurred
// 	assert.NoError(t, err)

// 	// Assert that the correct events were returned
// 	assert.Equal(t, expectedEvents, events)

// 	// Ensure all expectations were met
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

func TestEventRepository_GetAll(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	expectedEvents := []models.Event{
		{
			ID:        1,
			Name:      "Event 1",
			UserID:    1,
			StartTime: time.Date(2024, 10, 20, 10, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 10, 20, 12, 0, 0, 0, time.UTC),
			RoomID:    nil,
		},
		{
			ID:        2,
			Name:      "Event 2",
			UserID:    2,
			StartTime: time.Date(2024, 10, 21, 11, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 10, 21, 13, 0, 0, 0, time.UTC),
			RoomID:    func() *uint { r := uint(1); return &r }(),
		},
	}

	mock.ExpectQuery("SELECT \\* FROM `events`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "user_id", "start_time", "end_time", "room_id"}).
			AddRow(expectedEvents[0].ID, expectedEvents[0].Name, expectedEvents[0].UserID, expectedEvents[0].StartTime, expectedEvents[0].EndTime, expectedEvents[0].RoomID).
			AddRow(expectedEvents[1].ID, expectedEvents[1].Name, expectedEvents[1].UserID, expectedEvents[1].StartTime, expectedEvents[1].EndTime, expectedEvents[1].RoomID))

	repo := EventRepository{}

	events, err := repo.GetAll()

	assert.NoError(t, err)

	assert.Equal(t, expectedEvents, events)

	assert.NoError(t, mock.ExpectationsWereMet())
}

// func TestEventRepository_GetByID(t *testing.T) {
// 	// Set up the mock DB and sqlmock
// 	db, mock, err := config.SetupMockDB()
// 	assert.NoError(t, err)
// 	config.SetDB(db)

// 	// Event ID to search for
// 	eventID := uint(1)

// 	// Expected event to be returned
// 	expectedEvent := &models.Event{
// 		ID:        eventID,
// 		Name:      "Test Event",
// 		UserID:    1,
// 		StartTime: time.Date(2024, 10, 20, 10, 0, 0, 0, time.UTC),
// 		EndTime:   time.Date(2024, 10, 20, 12, 0, 0, 0, time.UTC),
// 		RoomID:    nil,
// 	}

// 	// Modify the expected query to match the SQL generated by GORM, with `LIMIT 1`
// 	mock.ExpectQuery("SELECT \\* FROM `events` WHERE `events`.`id` = ? ORDER BY `events`.`id` LIMIT 1").
// 		WithArgs(eventID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "user_id", "start_time", "end_time", "room_id"}).
// 			AddRow(expectedEvent.ID, expectedEvent.Name, expectedEvent.UserID, expectedEvent.StartTime, expectedEvent.EndTime, expectedEvent.RoomID))

// 	// Create the repository instance
// 	repo := EventRepository{}

// 	// Call the GetByID method
// 	event, err := repo.GetByID(eventID)

// 	// Assert that no errors occurred
// 	assert.NoError(t, err)

// 	// Assert that the correct event was returned
// 	assert.Equal(t, expectedEvent, event)

// 	// Ensure all expectations were met
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }






