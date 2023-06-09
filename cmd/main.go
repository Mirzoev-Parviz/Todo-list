package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	server "test"
	"test/config"
	"test/model"
	"test/pkg/handler"
	"test/pkg/repository"
	"test/pkg/service"

	_ "github.com/lib/pq"
)

func main() {

	config.DB.AutoMigrate(&model.User{}, &model.TODO{})
	db := config.ConnectDB()
	defer config.DisconnectDB(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error server")
		}
	}()
	fmt.Println("Server is runned...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Server is shutting down...")
	if err := srv.Sutdown(context.Background()); err != nil {
		fmt.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
