package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"step/handlers"
	"step/pkg/db"
	"step/repository"
	"step/service"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("failed to load dotenv: %w", err)
		return
	}

	cfg := db.GetConfig()

	database, err := db.GormInit(cfg)
	if err != nil {
		fmt.Println("failed to init db: %w", err)
		return
	}

	repo := repository.NewRepository(database)
	service := service.NewUsersService(repo)
	handler := handlers.NewHandler(service)
	router := handler.InitRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server failed to start: %w", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: %w", err)
	}

	fmt.Println("Server exited")
}
