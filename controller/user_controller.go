package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer func() {
		if ctx.Err() != context.Canceled {
			fmt.Println("ctx.Err():", ctx.Err())
		}
		cancel()
	}()
	// fmt.Println("ctx:", ctx)
	// time.Sleep(10 * time.Second)

	controller.UserService.Login(ctx, "usrname", "password")
	fmt.Fprint(w, "keren")
}
