package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type subscribeRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

func newSubscribeRepository(db *sql.DB, log *logrus.Logger) *subscribeRepository {
	return &subscribeRepository{
		db:  db,
		log: log,
	}
}

func (s *subscribeRepository) Create(subscription *models.Subscribe) (int, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "Create",
	}).Info("Create subscription")

	var subID int

	tx, err := s.db.Begin()
	if err != nil {
		s.log.Fatalf("Failed to start trsansaction, error: %v", err)
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
		s.log.Fatalf("Failed to scan values, error: %v", err)
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		s.log.Fatalf("Failed to commit transactio, error: %v", err)
		return 0, err
	}

	return subID, nil
}

func (s *subscribeRepository) GetAllByServiceName(serviceName string) ([]models.Subscribe, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "GetAllByServiceName",
	}).Info("Getting subscriptions by name")

	var subscribeList []models.Subscribe

	rows, err := s.db.Query(
		"SELECT id, service_name, price, user_id, start_date, end_date FROM subscribes WHERE service_name = $1",
		serviceName,
	)
	if err != nil {
		s.log.Fatalf("Failed to make query, error: %v", err)
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
			s.log.Fatalf("Failed to iterate thru rows, error: %v", err)
			return nil, err
		}
		subscribeList = append(subscribeList, subscribe)
	}
	if err = rows.Err(); err != nil {
		s.log.Fatalf("Failed to rows.Err(), error: %v", err)
		return nil, err
	}

	return subscribeList, nil
}

func (s *subscribeRepository) GetSubscriptionByID(subscriptionID int) (*models.Subscribe, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "GetSubscriptionByID",
	}).Info("Getting subscriptions by userID")

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
			s.log.Debugf("subscription with id %d not found", subscriptionID)
			return &subscription, fmt.Errorf("subscription with id %d not found", subscriptionID)
		}
		s.log.Fatalf("Failed to make query, error: %v", err)
		return &subscription, err
	}
	return &subscription, nil
}

func (s *subscribeRepository) GetUsersSubscriptions(userID uuid.UUID) ([]models.Subscribe, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "GetUsersSubscriptions",
	}).Info("Getting users subscriptions")

	var subscriptionList []models.Subscribe

	rows, err := s.db.Query(
		"SELECT id, service_name, price, user_id, start_date, end_date FROM subscribes WHERE user_id = $1",
		userID,
	)
	if err != nil {
		s.log.Fatalf("Failed to make query")
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
			s.log.Fatalf("Failed to scan row, error: %v", err)
			return nil, err
		}
		subscriptionList = append(subscriptionList, subscribe)
	}
	if err = rows.Err(); err != nil {
		s.log.Fatalf("Failed to rows.Err(), error: %v", err)
		return nil, err
	}

	return subscriptionList, nil
}

func (s *subscribeRepository) UpdateSubscription(subscriptionID, price int) (int, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "UpdateSubscription",
	}).Info("Updating subscription")

	sqlStatement := `UPDATE subscribes SET price = $1 WHERE id = $2;`
	res, err := s.db.Exec(sqlStatement, price, subscriptionID)
	if err != nil {
		s.log.Fatalf("Failed to make query, error: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.log.Debugf("Failed to es.RowsAffected(), error: %v", err)
		return 0, err
	}

	return int(rowsAffected), nil
}

func (s *subscribeRepository) DeleteSubscription(subscriptionID int) (int, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "DeleteSubscription",
	}).Info("Deleting subscription")

	tx, err := s.db.Begin()
	if err != nil {
		s.log.Fatalf("Failed to start transaction, error: %v", err)
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	res, err := tx.Exec("DELETE FROM subscribes WHERE id = $1", subscriptionID)
	if err != nil {
		s.log.Fatalf("Failed to make query, error: %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.log.Debugf("Failed to .RowsAffected(), error: %v", err)
		return 0, err
	}
	if rowsAffected == 0 {
		s.log.Infof("subscription with id %d not found", subscriptionID)
		return 0, fmt.Errorf("subscription with id %d not found", subscriptionID)
	}

	err = tx.Commit()
	if err != nil {
		s.log.Fatalf("Failed to end transaction, error: %v", err)
		return 0, err
	}

	return int(rowsAffected), nil
}

func (s *subscribeRepository) GetSubscriptionsPriceSum(startDate, endDate time.Time) ([]models.SubscriptionSummary, error) {
	s.log.WithFields(logrus.Fields{
		"layer":  "repository",
		"method": "GetSubscriptionsPriceSum",
	}).Info("Getting subscription sum of prices")

	query := `
	SELECT
		SUM(price),
		user_id,
		service_name
	FROM subscribes
	WHERE start_date <= $1 AND end_date >= $2
	GROUP BY
		user_id,
		service_name;
	`
	rows, err := s.db.Query(query, endDate, startDate)
	if err != nil {
		s.log.Fatalf("Failed to make query, error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.SubscriptionSummary
	for rows.Next() {
		var summary models.SubscriptionSummary
		err := rows.Scan(
			&summary.TotalPrice,
			&summary.UserID,
			&summary.ServiceName,
		)
		if err != nil {
			s.log.Fatalf("Failed to make query, error: %v", err)
			return nil, err
		}
		results = append(results, summary)
	}

	return results, nil
}
