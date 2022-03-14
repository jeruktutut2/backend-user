package route

import (
	"github.com/jeruktutut2/backend-user/controller"
	"github.com/jeruktutut2/backend-user/middleware"
	"github.com/julienschmidt/httprouter"
)

func UserRoute(router *httprouter.Router, userController controller.UserController) {
	router.POST("/api/v1/login", middleware.MultipleMiddleware(userController.Login, middleware.Middleware1, middleware.Middleware2))
}
