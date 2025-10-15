package handlers

import (
	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
	log     *logrus.Logger
}

func NewHandler(service *service.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
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
			subscribes.GET("/service/:serviceName", h.GetAllSubscriptionsByServiceName)
			subscribes.GET("/:serviceID", h.GetSubscriptionByID)
			subscribes.GET("/user_subscriptions/:userID", h.GetUsersSubscriptions)
			subscribes.PUT("/update_subscription/:subscriptionID/:price", h.UpdateSubscription)
			subscribes.DELETE("/delete_subscription/:subscriptionID", h.DeleteSubscription)
			subscribes.GET("/list/:startDate/:endDate", h.GetSubscriptionsPriceSum)
		}
	}
	return router
}
