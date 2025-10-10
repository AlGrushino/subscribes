package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       int
	Name     string
	Password string
}

type Subscribe struct {
	Id          int
	ServiceName string
	Price       int
	UserId      int
	UserUUID    uuid.UUID
	StartDate   time.Time
	EndDate     time.Time
}

// удалить перед релизом
type Person struct {
	Name string
	Age  int
}
