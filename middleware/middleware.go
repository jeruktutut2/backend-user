package middleware

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

func MultipleMiddleware(h httprouter.Handle, middlewares ...Middleware) httprouter.Handle {
	if len(middlewares) < 1 {
		return h
	}

	wrapped := h

	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped
}

func Middleware1(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("Middleware1")
		// c := 1
		// if c == 1 {
		// 	fmt.Fprint(w, "Middleware1 c 1")
		// } else {
		// 	next(w, r, params)
		// }

		next(w, r, params)
	}
}

func Middleware2(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("Middleware2")
		// c := 1
		// if c == 1 {
		// 	fmt.Fprint(w, "Middleware2 c")
		// } else {
		// 	next(w, r, params)
		// }
		next(w, r, params)
	}
}
