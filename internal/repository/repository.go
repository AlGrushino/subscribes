package repository

import (
	"database/sql"

	"github.com/AlGrushino/subscribes/internal/repository/models"
)

type Subscribe interface {
	Create(subscription *models.Subscribe) (int, error)
	GetAllByServiceName(serviceName string) ([]models.Subscribe, error)
	GetSubscriptionByID(subscriptionID int) (*models.Subscribe, error)
}

type Repository struct {
	Subscribe
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Subscribe: newSubscribeRepository(db),
	}
}
