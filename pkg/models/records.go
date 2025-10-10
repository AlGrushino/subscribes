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
	Id          int
	ServiceName string
	Price       int
	UserUUID    uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}
