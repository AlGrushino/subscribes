package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateSubscription(c *gin.Context) {
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "CreateSubscription",
	}).Info("Creating subscription")

	subscribe := &models.Subscribe{}
	subscriptionID, err := h.service.Create(c, subscribe)
	if err != nil {
		h.log.Fatal("Failed to create subscription")

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      subscriptionID,
		"message": "Subscription created successfully",
	})
}

func (h *Handler) GetAllSubscriptionsByServiceName(c *gin.Context) {
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "GetAllSubscriptionsByServiceName",
	}).Info("Getting all subs")

	serviceName := c.Param("serviceName")
	if serviceName == "" {
		h.log.Fatalf("service name is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "serviceName parameter is required"},
		)
		return
	}

	list, err := h.service.GetAllByServiceName(serviceName)
	if err != nil {
		h.log.Fatalf("Failed to call h.service.GetAllByServiceName(serviceName), error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to get subscriptions"},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"subscriptionsList": list,
			"count":             len(list),
		},
	)
}

func (h *Handler) GetSubscriptionByID(c *gin.Context) {
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "GetSubscriptionByID",
	}).Info("Getting subscription by ID")

	serviceID := c.Param("serviceID")
	if serviceID == "" {
		h.log.Fatal("serviceID param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "serviceID parameter is required"})
		return
	}

	serviceIDInt, err := strconv.Atoi(serviceID)
	if err != nil {
		h.log.Fatalf("Failed to Atoi, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert"})
		return
	}

	subscription, err := h.service.Subscribe.GetSubscriptionByID(serviceIDInt)
	if err != nil {
		h.log.Fatalf("Failed to Atoi, error: %v", err)

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
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "GetUsersSubscriptions",
	}).Info("Getting users subscriptions")

	userID := c.Param("userID")
	if userID == "" {
		h.log.Fatal("userID param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "userID parameter is required"})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		h.log.Fatalf("Failed to parse userID, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert"})
		return
	}

	subscriptions, err := h.service.GetUsersSubscriptions(userUUID)
	if err != nil {
		h.log.Fatalf("Failed to GetUsersSubscriptions(userUUID), error: %v", err)

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
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "UpdateSubscription",
	}).Info("Updating subscription")

	subscriptionID := c.Param("subscriptionID")
	price := c.Param("price")

	if subscriptionID == "" {
		h.log.Fatal("subscriptionID param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "subscriptionID parameter is required"})
		return
	}
	if price == "" {
		h.log.Fatal("price param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "price parameter is required"})
		return
	}

	subscriptionIDInt, err := strconv.Atoi(subscriptionID)
	if err != nil {
		h.log.Fatalf("failed to atoi subscriptionIDInt, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert subscriptionID"})
		return
	}

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		h.log.Fatalf("failed to atoi price, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert price"})
		return
	}

	rowsAffected, err := h.service.UpdateSubscription(subscriptionIDInt, priceInt)
	if err != nil {
		h.log.Fatalf("failed to UpdateSubscription(subscriptionIDInt, priceInt), error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to udpate subscription"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"rowsAffected": rowsAffected})
}

func (h *Handler) DeleteSubscription(c *gin.Context) {
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "DeleteSubscription",
	}).Info("Delete subscription")

	subscriptionID := c.Param("subscriptionID")
	if subscriptionID == "" {
		h.log.Fatalf("subscriptionID param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "subscriptionID parameter is required"})
		return
	}

	subscriptionIDInt, err := strconv.Atoi(subscriptionID)
	if err != nil {
		h.log.Fatalf("failed to atoi, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert subscriptionID"})
		return
	}

	rowsAffected, err := h.service.DeleteSubscription(subscriptionIDInt)
	if err != nil {
		h.log.Fatalf("failed to DeleteSubscription(subscriptionIDInt), error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "faied to delete subscription"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"rowsAffected": rowsAffected})
}

func (h *Handler) GetSubscriptionsPriceSum(c *gin.Context) {
	h.log.WithFields(logrus.Fields{
		"layer":  "handler",
		"method": "GetSubscriptionsPriceSum",
	}).Info("Getting subscription price and sum")

	startDate := c.Param("startDate")
	if startDate == "" {
		h.log.Fatalf("startDate param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "startDate parameter is required"},
		)
		return
	}

	endDate := c.Param("endDate")
	if endDate == "" {
		h.log.Fatalf("endDate param is empty")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "endDate parameter is required"},
		)
		return
	}

	layOut := "01-2006"
	parsedStartDate, err := time.Parse(layOut, startDate)
	if err != nil {
		h.log.Fatalf("failed to parse startDate, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert startDate"},
		)
		return
	}

	parsedEndDate, err := time.Parse(layOut, endDate)
	if err != nil {
		h.log.Fatalf("failed to parse endDate, error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to convert endDate"},
		)
		return
	}

	subscriptionList, err := h.service.GetSubscriptionsPriceSum(parsedStartDate, parsedEndDate)
	if err != nil {
		h.log.Fatalf("failed to GetSubscriptionsPriceSum(parsedStartDate, parsedEndDate), error: %v", err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "failed to find records"},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"subscriptionList": subscriptionList,
			"count":            len(subscriptionList),
		},
	)
}
