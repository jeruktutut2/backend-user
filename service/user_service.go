package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jeruktutut2/backend-user/exception"
	"github.com/jeruktutut2/backend-user/repository"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (err error)
}

type UserServiceImplementation struct {
	DB             *sql.DB
	UserRepository repository.UserRepository
}

func NewUserService(DB *sql.DB, userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{
		DB:             DB,
		UserRepository: userRepository,
	}
}

func (service *UserServiceImplementation) Login(ctx context.Context, username string, password string) (err error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	exception.LogFatallnIfError(err)

	test1 := service.UserRepository.TestSleep(tx, ctx)
	rowsAffectedInsertTable1 := service.UserRepository.InsertTable1(tx, ctx)
	test1 = service.UserRepository.TestSleep(tx, ctx)
	rowsAffectedInsertTable1 = service.UserRepository.InsertTable1(tx, ctx)
	test1 = service.UserRepository.TestSleep(tx, ctx)
	fmt.Print("test", test1, rowsAffectedInsertTable1)
	return errors.New("error")
}
