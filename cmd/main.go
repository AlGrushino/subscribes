package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/AlGrushino/subscribes/internal/handlers"
	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/AlGrushino/subscribes/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func checkRoutes(router *gin.Engine) {
	routes := router.Routes()
	fmt.Printf("Registered routes:\n")
	for _, route := range routes {
		fmt.Printf("- %s %s\n", route.Method, route.Path)
	}
}

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

	// Добавьте эту проверку
	checkRoutes(router)

	fmt.Println("=== STARTING SERVER DIRECTLY ON :8080 ===")

	// ЗАКОММЕНТИРУЙТЕ ВЕСЬ server package код
	// server := server.NewServer()
	// go func() {
	// 	if err := server.Run("8080", router.Handler()); err != nil {
	// 		log.Printf("Server: %v", err)
	// 	}
	// }()

	// Запустите напрямую
	go func() {
		fmt.Println("Gin server starting on :8080...")
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Подождите немного и проверьте
	time.Sleep(3 * time.Second)

	// Сделайте тестовый запрос из самого приложения
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("=== MAKING TEST REQUEST ===")
		testRequest()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}

func testRequest() {
	client := &http.Client{Timeout: 5 * time.Second}

	jsonData := `{
		"service_name": "Sber Plus",
		"price": 400,
		"user_id": "fcd30c1d-fa2b-4d41-9512-c27c245494ec",
		"start_date": "07-2025"
	}`

	resp, err := client.Post("http://localhost:8080/api/subscribes",
		"application/json",
		strings.NewReader(jsonData))
	if err != nil {
		fmt.Printf("TEST REQUEST ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("TEST RESPONSE: Status=%d, Body=%s\n", resp.StatusCode, string(body))
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/AlGrushino/subscribes/internal/handlers"
// 	"github.com/AlGrushino/subscribes/internal/repository"
// 	"github.com/AlGrushino/subscribes/internal/service"
// 	"github.com/AlGrushino/subscribes/pkg/db"
// 	"github.com/joho/godotenv"
// )

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

// 	fmt.Println("Starting server on :8080...")

// 	// Создаем свой http.Server для graceful shutdown
// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: router,
// 	}

// 	// Запускаем сервер в горутине
// 	go func() {
// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("Server failed to start: %v", err)
// 		}
// 	}()

// 	fmt.Println("Server is running on http://localhost:8080")
// 	fmt.Println("Use Ctrl+C to stop the server")

// 	// Ожидаем сигнал завершения
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit

// 	log.Println("Shutting down server...")

// 	// Graceful shutdown
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := server.Shutdown(ctx); err != nil {
// 		log.Fatalf("Server forced to shutdown: %v", err)
// 	}

// 	log.Println("Server exited")
// }
