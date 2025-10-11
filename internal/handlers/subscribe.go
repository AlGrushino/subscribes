package handlers

import (
	"net/http"
	"time"

	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/gin-gonic/gin"
)

type CreateSubscribeRequest struct {
	ServiceName string  `json:"service_name" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	UserID      string  `json:"user_id" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     *string `json:"end_date,omitempty"`
}

func (h *Handler) CreateSubscription(c *gin.Context) {
	var req CreateSubscribeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
		return
	}

	var endDate *time.Time
	if req.EndDate != nil && *req.EndDate != "" {
		parsedEndDate, err := time.Parse("01-2006", *req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
			return
		}
		endDate = &parsedEndDate
	}

	parsedUserID, err := service.ParseUUID(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	subscribe := &models.Subscribe{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserUUID:    parsedUserID,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	subscriptionID, err := h.service.Subscribe.Create(subscribe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Did not work to add subscription"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      subscriptionID,
		"message": "Subscription created successfully",
	})
}

func (h *Handler) GetAllSubscriptionsByServiceName(c *gin.Context) {
	serviceName := c.Param("serviceName")
	if serviceName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "service_name parameter is required"})
		return
	}

	subscriptionList, err := h.service.Subscribe.GetAllByServiceName(serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscriptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"service_name":  serviceName,
		"subscriptions": subscriptionList,
		"count":         len(subscriptionList),
	})
}
