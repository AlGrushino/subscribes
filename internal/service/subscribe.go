package service

import (
	"time"

	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type subscribeService struct {
	repository repository.Subscribe
	log        *logrus.Logger
}

func newSubscribeService(repository repository.Subscribe, log *logrus.Logger) *subscribeService {
	return &subscribeService{
		repository: repository,
		log:        log,
	}
}

// CreateSubscribeRequest
type CreateSubscribeRequest struct {
	ServiceName string  `json:"service_name" binding:"required" exmaple:"Netflix"`
	Price       int     `json:"price" binding:"required" exmaple:"322"`
	UserID      string  `json:"user_id" binding:"required" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	StartDate   string  `json:"start_date" binding:"required" example:"03-2024"`
	EndDate     *string `json:"end_date,omitempty" example:"05-2024"`
}

func (s *subscribeService) Create(c *gin.Context, subscription *models.Subscribe) (int, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "Create",
	}).Info("Create subscription")

	var req CreateSubscribeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Fatalf("Failed to bind JSON, error: %v", err)
		return 0, err
	}

	if err := handleDates(subscription, req.StartDate, *req.EndDate); err != nil {
		s.log.Fatalf("Failed to handle dates, error: %v", err)
		return 0, err
	}

	if err := handlePrice(subscription, req.Price); err != nil {
		s.log.Fatalf("Failed to handle price, error: %v", err)
		return 0, err
	}

	if err := handleServiceName(subscription, req.ServiceName); err != nil {
		s.log.Fatalf("Failed to handle serviceName, error: %v", err)
		return 0, err
	}

	if err := handleUserID(subscription, req.UserID); err != nil {
		s.log.Fatalf("Failed to handle userID, error: %v", err)
		return 0, err
	}

	id, err := s.repository.Create(subscription)
	if err != nil {
		s.log.Fatalf("Failed to create subscription, error: %v", err)
		return 0, err
	}

	return id, nil
}

func (s *subscribeService) GetAllByServiceName(serviceName string) ([]models.Subscribe, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "GetAllByServiceName",
	}).Info("Getting all subs")

	return s.repository.GetAllByServiceName(serviceName)
}

func (s *subscribeService) GetSubscriptionByID(subscriptionID int) (*models.Subscribe, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "GetSubscriptionByID",
	}).Info("Getting subscription")

	return s.repository.GetSubscriptionByID(subscriptionID)
}

func (s *subscribeService) GetUsersSubscriptions(userID uuid.UUID) ([]models.Subscribe, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "GetUsersSubscriptions",
	}).Info("Getting users subs")

	return s.repository.GetUsersSubscriptions(userID)
}

func (s *subscribeService) UpdateSubscription(subscriptionID, price int) (int, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "UpdateSubscription",
	}).Info("Updating subscription")

	return s.repository.UpdateSubscription(subscriptionID, price)
}

func (s *subscribeService) DeleteSubscription(subscriptionID int) (int, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "DeleteSubscription",
	}).Info("Deleting subscription")

	return s.repository.DeleteSubscription(subscriptionID)
}

func (s *subscribeService) GetSubscriptionsPriceSum(startDate, endDate time.Time) ([]models.SubscriptionSummary, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "service",
		"method": "GetSubscriptionsPriceSum",
	}).Info("Getting subscription price cum")

	return s.repository.GetSubscriptionsPriceSum(startDate, endDate)
}
