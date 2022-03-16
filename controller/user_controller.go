package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImplementation struct {
}

func NewUserController() UserController {
	return &UserControllerImplementation{}
}

func (controller *UserControllerImplementation) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer func() {
		if ctx.Err() != nil {
			fmt.Println("ctx.Err():", ctx.Err())
		}
		cancel()
	}()
	fmt.Println("ctx:", ctx)
	// time.Sleep(10 * time.Second)
	fmt.Fprint(w, "keren")
}
