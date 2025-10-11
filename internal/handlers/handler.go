package handlers

import (
	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	// отладка
	// router.Use(func(c *gin.Context) {
	// 	fmt.Printf("Incoming request: %s %s\n", c.Request.Method, c.Request.URL.Path)
	// 	c.Next()
	// })
	// отладка

	api := router.Group("/api")
	{
		subscribes := api.Group("/subscribes")
		{
			subscribes.POST("", h.CreateSubscription)
			subscribes.GET("/:serviceName", h.GetAllSubscriptionsByServiceName)
		}
	}
	return router
}
