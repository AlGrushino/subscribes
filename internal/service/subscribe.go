package service

import (
	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
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

func (s *subscribeService) Create(subscription *models.Subscribe) (int, error) {
	return s.repository.Create(subscription)
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
