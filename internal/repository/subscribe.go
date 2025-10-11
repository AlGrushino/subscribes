package repository

import (
	"database/sql"
	"fmt"

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
		return 0, err
	}

	return subID, nil
}

func (s *subscribeRepository) GetAllByServiceName(serviceName string) ([]models.Subscribe, error) {
	var subscribeList []models.Subscribe

	rows, err := s.db.Query(
		"SELECT id, service_name, price, user_id, start_date, end_date FROM subscribes WHERE service_name = $1",
		serviceName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subscribe models.Subscribe
		err := rows.Scan(
			&subscribe.Id,
			&subscribe.ServiceName,
			&subscribe.Price,
			&subscribe.UserUUID,
			&subscribe.StartDate,
			&subscribe.EndDate,
		)
		if err != nil {
			return nil, err
		}
		subscribeList = append(subscribeList, subscribe)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subscribeList, nil
}

func (s *subscribeRepository) GetSubscriptionByID(subscriptionID int) (*models.Subscribe, error) {
	var subscription models.Subscribe

	err := s.db.QueryRow(
		"SELECT id, service_name, price, user_id, start_date, end_date FROM subscribes WHERE id = $1",
		subscriptionID,
	).Scan(
		&subscription.Id,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserUUID,
		&subscription.StartDate,
		&subscription.EndDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &subscription, fmt.Errorf("subscription with id %d not found", subscriptionID)
		}
		return &subscription, err
	}
	return &subscription, nil
}
