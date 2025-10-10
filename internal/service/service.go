package service

import (
	"github.com/AlGrushino/subscribes/internal/repository"
	"github.com/AlGrushino/subscribes/internal/repository/models"
)

type Service struct {
	Subscribe Subscribe
}

type Subscribe interface {
	Create(subscription *models.Subscribe) (int, error)
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Subscribe: newSubscribeService(repository.Subscribe),
	}
}
