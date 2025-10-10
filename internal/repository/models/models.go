package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Password string
	Email    string
}

type Subscribe struct {
	Id          int        `db:"id"`
	ServiceName string     `db:"service_name"`
	Price       int        `db:"price"`
	UserUUID    uuid.UUID  `db:"user_id"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     *time.Time `db:"end_date"`
}
