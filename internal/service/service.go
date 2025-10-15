package service

import (
	"time"

	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	log *logrus.Logger
}

func NewService(repository *repository.Repository, log *logrus.Logger) *Service {
	return &Service{
		Subscribe: newSubscribeService(repository.Subscribe, log),
		log:       log,
	}
}
