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

	"github.com/AlGrushino/subscribes/internal/handlers"
	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/AlGrushino/subscribes/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := db.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	database, err := db.DBInit(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer database.Close()

	repo := repository.NewRepository(database)
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)
	router := handler.InitRoutes()

	fmt.Println("Starting server on :8080...")

	// Создаем свой http.Server для graceful shutdown
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Запускаем сервер в горутине
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	fmt.Println("Server is running on http://localhost:8080")
	fmt.Println("Use Ctrl+C to stop the server")

	// Ожидаем сигнал завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
