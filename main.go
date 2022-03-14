package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jeruktutut2/backend-user/controller"
	"github.com/jeruktutut2/backend-user/route"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("tes")

	router := httprouter.New()

	userController := controller.NewUserController()
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
	log.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	log.Println("Server Exited Properly")
}
