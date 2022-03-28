package mock

import (
	"context"
	"database/sql"

	"github.com/jeruktutut2/backend-user/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (userRepositoryMock *UserRepositoryMock) GetUserByUsername(tx *sql.Tx, ctx context.Context, username string) (user entity.User) {
	arguments := userRepositoryMock.Mock.Called(tx, ctx, username)
	return arguments.Get(0).(entity.User)
}
