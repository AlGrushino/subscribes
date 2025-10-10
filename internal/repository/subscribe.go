package repository

import (
	"database/sql"

	"github.com/AlGrushino/subscribes/internal/repository/models"
)

type subscribeRepository struct {
	db *sql.DB
}

func newSubscribeRepository(db *sql.DB) *subscribeRepository {
	return &subscribeRepository{
		db: db,
	}
}

func (s *subscribeRepository) Create(subscription *models.Subscribe) (int, error) {
	var subID int

	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	row := tx.QueryRow(
		"INSERT INTO subscribes (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		subscription.ServiceName,
		subscription.Price,
		subscription.UserUUID,
		subscription.StartDate,
		subscription.EndDate,
	)
	err = row.Scan(&subID)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, nil
	}

	return subID, nil
}
