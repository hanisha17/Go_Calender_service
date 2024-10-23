package repositories

import (
	"calender-service/config"
	"calender-service/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := config.SetupMockDB()
	assert.NoError(t, err)
	config.SetDB(db)

	user := &models.User{
		ID:    1,
		Name:  "John Doe",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.Name, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := UserRepository{}
	err = repo.Create(user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

}


// func TestUserRepository_GetByID(t *testing.T) {
//     db, mock, err := config.SetupMockDB()
//     assert.NoError(t, err)
//     config.SetDB(db)

//     expectedUser := &models.User{
//         ID:   1,
//         Name: "John Doe",
//     }

//     mock.ExpectQuery("SELECT \\* FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT \\?").
//         WithArgs(expectedUser.ID).
//         WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//             AddRow(expectedUser.ID, expectedUser.Name))

//     repo := UserRepository{}
//     user, err := repo.GetByID(1)

//     assert.NoError(t, err)
//     assert.Equal(t, expectedUser, user)
//     assert.NoError(t, mock.ExpectationsWereMet())
// }

