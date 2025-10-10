package handlers

import (
	"database/sql"
	"fmt"

	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/google/uuid"
)

func GetUserSubscribes(user_id uuid.UUID, db *sql.DB) ([]models.Subscribe, error) {
	subscribes := []models.Subscribe{}

	rows, err := db.Query(
		"SELECT id, service_name, price, start_date, end_date FROM subscribes WHERE user_id = $1",
		user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subscribe models.Subscribe
		err := rows.Scan(&subscribe.Id,
			&subscribe.ServiceName,
			&subscribe.Price,
			&subscribe.StartDate,
			&subscribe.EndDate)
		if err != nil {
			return nil, err
		}
		subscribes = append(subscribes, subscribe)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscribes, nil
}

func GetFirstUserId(db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	row := db.QueryRow("SELECT id FROM users LIMIT 1")
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return id, fmt.Errorf("no users")
		}
		return id, err
	}
	return id, nil
}
