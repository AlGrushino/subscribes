package service

import (
	"time"

	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type subscribeService struct {
	repository repository.Subscribe
}

func newSubscribeService(repository repository.Subscribe) *subscribeService {
	return &subscribeService{
		repository: repository,
	}
}

type CreateSubscribeRequest struct {
	ServiceName string  `json:"service_name" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	UserID      string  `json:"user_id" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     *string `json:"end_date,omitempty"`
}

func (s *subscribeService) Create(c *gin.Context, subscription *models.Subscribe) (int, error) {
	var req CreateSubscribeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return 0, err
	}

	if err := handleDates(subscription, req.StartDate, *req.EndDate); err != nil {
		return 0, err
	}

	if err := handlePrice(subscription, req.Price); err != nil {
		return 0, err
	}

	if err := handleServiceName(subscription, req.ServiceName); err != nil {
		return 0, err
	}

	if err := handleUserID(subscription, req.UserID); err != nil {
		return 0, err
	}

	id, err := s.repository.Create(subscription)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *subscribeService) GetAllByServiceName(serviceName string) ([]models.Subscribe, error) {
	return s.repository.GetAllByServiceName(serviceName)
}

func (s *subscribeService) GetSubscriptionByID(subscriptionID int) (*models.Subscribe, error) {
	return s.repository.GetSubscriptionByID(subscriptionID)
}

func (s *subscribeService) GetUsersSubscriptions(userID uuid.UUID) ([]models.Subscribe, error) {
	return s.repository.GetUsersSubscriptions(userID)
}

func (s *subscribeService) UpdateSubscription(subscriptionID, price int) (int, error) {
	return s.repository.UpdateSubscription(subscriptionID, price)
}

func (s *subscribeService) DeleteSubscription(subscriptionID int) (int, error) {
	return s.repository.DeleteSubscription(subscriptionID)
}

// GetSubscriptionsPriceSum(startDate, endDate time.Time) ([]SubscriptionSummary, error)
func (s *subscribeService) GetSubscriptionsPriceSum(startDate, endDate time.Time) ([]models.SubscriptionSummary, error) {
	return s.repository.GetSubscriptionsPriceSum(startDate, endDate)
}
