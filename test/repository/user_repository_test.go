package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jeruktutut2/backend-user/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	username := "username"

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, username, password FROM user WHERE username = ?").
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow("id", "username", "password"))
	mock.ExpectCommit()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal("err:", err)
	}
	userRepository := repository.NewUserRepository()
	user := userRepository.GetUserByUsername(tx, ctx, username)
	assert.Equal(t, user.Id, "id")
	assert.Equal(t, user.Username, username)
}
