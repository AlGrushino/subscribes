package service

import (
	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
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
