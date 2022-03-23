package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/jeruktutut2/backend-user/exception"
	"github.com/jeruktutut2/backend-user/model/request"
	"github.com/jeruktutut2/backend-user/model/response"
	"github.com/jeruktutut2/backend-user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, userLoginRequest request.UserLoginRequest) (userLoginResponse response.UserLoginResponse)
}

type UserServiceImplementation struct {
	DB             *sql.DB
	Validate       *validator.Validate
	UserRepository repository.UserRepository
}

func NewUserService(DB *sql.DB, validate *validator.Validate, userRepository repository.UserRepository) UserService {
	return &UserServiceImplementation{
		DB:             DB,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (service *UserServiceImplementation) Login(ctx context.Context, userLoginRequest request.UserLoginRequest) (userLoginResponse response.UserLoginResponse) {
	err := service.Validate.Struct(userLoginRequest)
	exception.PanicIfErrorValidator(err)
	tx, err := service.DB.BeginTx(ctx, nil)
	exception.PanicIfError(err)
	user := service.UserRepository.GetUserByUsername(tx, ctx, userLoginRequest.Username)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	exception.PanicIfErrorAndRollback(err, tx)
	err = tx.Commit()
	exception.PanicIfErrorAndRollback(err, tx)
	userLoginResponse = response.ToUserLoginResponse(user.Id, user.Username)
	return userLoginResponse
}
