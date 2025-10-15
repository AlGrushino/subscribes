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
	log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "NewHandler",
	}).Info("Create new handler")

	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "InitRoutes",
	}).Info("Initing routes")

	router := gin.Default()

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
