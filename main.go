package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jeruktutut2/backend-user/configuration"
	"github.com/jeruktutut2/backend-user/controller"
	"github.com/jeruktutut2/backend-user/exception"
	"github.com/jeruktutut2/backend-user/repository"
	"github.com/jeruktutut2/backend-user/route"
	"github.com/jeruktutut2/backend-user/service"
	"github.com/jeruktutut2/backend-user/util"
	"github.com/julienschmidt/httprouter"
)

func main() {
	configuration := configuration.NewConfiguration()
	// fmt.Println("configuration:", configuration)
	util.SetTimezone(configuration.Timezone.Timezone)
	databaseConnection := util.NewDatabaseConnection(configuration.Database)
	// fmt.Println("databaseConnection:", databaseConnection)
	validator := validator.New()
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(databaseConnection, validator, userRepository)
	userController := controller.NewUserController(userService)
	route.UserRoute(router, userController)

	server := &http.Server{
		Addr:    ":10001",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server Started")

	<-done

	util.Close(databaseConnection)
	log.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	log.Println("Server Exited Properly")
}
