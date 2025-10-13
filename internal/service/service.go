package service

import (
	"time"

	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Subscribe interface {
	Create(c *gin.Context, subscription *models.Subscribe) (int, error)
	GetAllByServiceName(serviceName string) ([]models.Subscribe, error)
	GetSubscriptionByID(subscriptionID int) (*models.Subscribe, error)
	GetUsersSubscriptions(userID uuid.UUID) ([]models.Subscribe, error)
	UpdateSubscription(subscriptionID, price int) (int, error)
	DeleteSubscription(subscriptionID int) (int, error)
	GetSubscriptionsPriceSum(startDate, endDate time.Time) ([]models.SubscriptionSummary, error)
}

type Service struct {
	Subscribe
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Subscribe: newSubscribeService(repository.Subscribe),
	}
}
