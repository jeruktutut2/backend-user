package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jeruktutut2/backend-user/entity"
	"github.com/jeruktutut2/backend-user/exception"
)

type UserRepository interface {
	TestSleep(tx *sql.Tx, ctx context.Context) (test string)
	InsertTable1(tx *sql.Tx, ctx context.Context) (rowsAffected int64)
	GetUserByUsername(tx *sql.Tx, ctx context.Context, username string) (user entity.User)
}

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}

func (repository *UserRepositoryImplementation) TestSleep(tx *sql.Tx, ctx context.Context) (test string) {
	sql := "SELECT SLEEP(10) AS sleep"
	rows, err := tx.QueryContext(ctx, sql)
	exception.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&test)
		exception.LogFatallnIfError(err)
	}
	return test
}

func (repository *UserRepositoryImplementation) InsertTable1(tx *sql.Tx, ctx context.Context) (rowsAffected int64) {
	sql := "INSERT INTO golang_example.table_1 (id, name) VALUES(?, ?);"
	result, err := tx.ExecContext(ctx, sql, uuid.NewString(), "name 1")
	exception.PanicIfError(err)
	rowsAffected, err = result.RowsAffected()
	exception.LogFatallnIfError(err)
	return rowsAffected
}

func (repository *UserRepositoryImplementation) GetUserByUsername(tx *sql.Tx, ctx context.Context, username string) (user entity.User) {
	query := "SELECT id, username, password FROM user WHERE username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	exception.PanicIfErrorAndRollback(err, tx)
	defer rows.Close()

	user = entity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		exception.PanicIfErrorAndRollback(err, tx)
	}
	return user
}
