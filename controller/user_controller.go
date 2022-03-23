package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jeruktutut2/backend-user/model/request"
	"github.com/jeruktutut2/backend-user/model/response"
	"github.com/jeruktutut2/backend-user/service"
	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImplementation{
		UserService: userService,
	}
}

func (controller *UserControllerImplementation) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer func() {
		response.ResponseHttpContext(w, ctx.Err())
		cancel()
	}()
	userLoginRequest := request.UserLoginRequest{}
	request.ReadFromRequestBody(r, &userLoginRequest)
	fmt.Println("userLoginRequest:", userLoginRequest)
	userLoginResponse := controller.UserService.Login(ctx, userLoginRequest)
	response.ResponseHttp(w, http.StatusOK, userLoginResponse, nil)
}
