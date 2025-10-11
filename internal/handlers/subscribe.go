package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/AlGrushino/subscribes/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "service_name parameter is required"})
		return
	}

	subscriptionList, err := h.service.Subscribe.GetAllByServiceName(serviceName)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to get subscriptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"service_name":  serviceName,
		"subscriptions": subscriptionList,
		"count":         len(subscriptionList),
	})
}

func (h *Handler) GetSubscriptionByID(c *gin.Context) {
	serviceID := c.Param("serviceID")
	if serviceID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "serviceID parameter is required"})
		return
	}

	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert"})
		return
	}

	subscription, err := h.service.Subscribe.GetSubscriptionByID(serviceIDInt)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to get subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subscription": *subscription,
	})
}

func (h *Handler) GetUsersSubscriptions(c *gin.Context) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "userID parameter is required"})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert"})
		return
	}

	subscriptions, err := h.service.GetUsersSubscriptions(userUUID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to get subscriptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subscriptions": subscriptions,
		"count":         len(subscriptions),
	})
}

func (h *Handler) UpdateSubscription(c *gin.Context) {
	subscriptionID := c.Param("subscriptionID")
	price := c.Param("price")

	if subscriptionID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "subscriptionID parameter is required"})
		return
	}
	if price == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "price parameter is required"})
		return
	}

	subscriptionIDInt, err := strconv.Atoi(subscriptionID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert subscriptionID"})
		return
	}

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert price"})
		return
	}

	rowsAffected, err := h.service.UpdateSubscription(subscriptionIDInt, priceInt)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to udpate subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rowsAffected": rowsAffected})
}
