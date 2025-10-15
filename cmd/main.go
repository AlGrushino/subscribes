package main

import (
	"context"
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
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	logFile, err := os.OpenFile("../logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg, err := db.GetConfig(log)
	if err != nil {
		log.Fatal(err)
		return
	}

	database, err := db.DBInit(log, cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer database.Close()

	repo := repository.NewRepository(database)
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)
	router := handler.InitRoutes()

	// fmt.Println("Starting server on :8080...")
	log.Info("Server started")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// fmt.Println("Server is running on http://localhost:8080")
	// fmt.Println("Use Ctrl+C to stop the server")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
