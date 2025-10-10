// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/AlGrushino/subscribes/internal/handlers"
// 	"github.com/AlGrushino/subscribes/internal/repository"
// 	"github.com/AlGrushino/subscribes/internal/service"
// 	"github.com/AlGrushino/subscribes/pkg/db"
// 	"github.com/AlGrushino/subscribes/pkg/server"
// 	"github.com/gin-gonic/gin"
// 	_ "github.com/golang-migrate/migrate/v4/database/postgres"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"

// 	"github.com/joho/godotenv"
// )

// func checkRoutes(router *gin.Engine) {
// 	routes := router.Routes()
// 	fmt.Printf("Registered routes:\n")
// 	for _, route := range routes {
// 		fmt.Printf("- %s %s\n", route.Method, route.Path)
// 	}
// }

// func main() {
// 	err := godotenv.Load("../.env")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	cfg, err := db.GetConfig()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	database, err := db.DBInit(cfg)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer database.Close()

// 	repo := repository.NewRepository(database)
// 	serv := service.NewService(repo)
// 	handler := handlers.NewHandler(serv)
// 	router := handler.InitRoutes()

// 	checkRoutes(router)

// 	server := server.NewServer()

// 	go func() {
// 		if err := server.Run("8080", router.Handler()); err != nil {
// 			log.Printf("Server: %v", err)
// 		}
// 	}()

// 	// m, err := migrate.New(
// 	// 	"file://../migrations",
// 	// 	db.GetConnStr(cfg))
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// defer m.Close()

// 	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// userID, err := uuid.Parse("fcd30c1d-fa2b-4d41-9512-c27c245494ec")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// 	// log.Fatal("Invalid UUID:", err)
// 	// }

// 	// userID, err := handlers.GetFirstUserId(database)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// subscribes, err := handlers.GetUserSubscribes(userID, database)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// fmt.Printf("Найдено %d подписок для пользователя %s:\n", len(subscribes), userID)
// 	// for _, sub := range subscribes {
// 	// 	endDateStr := "бессрочная"
// 	// 	if sub.EndDate != nil {
// 	// 		endDateStr = sub.EndDate.Format("2006-01-02")
// 	// 	}
// 	// 	fmt.Printf("- %s: %d руб. (с %s по %s)\n",
// 	// 		sub.ServiceName,
// 	// 		sub.Price,
// 	// 		sub.StartDate.Format("2006-01-02"),
// 	// 		endDateStr)
// 	// }

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit

// 	log.Println("Shutting down server...")

// 	// Graceful shutdown с таймаутом
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := server.Shutdown(ctx); err != nil {
// 		log.Fatalf("Server forced to shutdown: %v", err)
// 	}

// 	log.Println("Server exited")
// }
