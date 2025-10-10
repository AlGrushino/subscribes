package handlers

import (
	"fmt"

	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: *service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	// отладка
	router.Use(func(c *gin.Context) {
		fmt.Printf("Incoming request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
	// отладка

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		// на случай, если будет авторизация
		// AllowCredentials: true,
		// MaxAge:           12 * 3600,
	}))

	api := router.Group("/api")
	{
		subscribes := api.Group("/subscribes")
		{
			subscribes.POST("", h.CreateSubscription)
		}
	}
	return router
}
