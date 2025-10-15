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

// Subscribe represents a subscription entity in the system
// @Description Subscription model containing all subscription details
type Subscribe struct {
	Id          int        `db:"id" example:"1"`
	ServiceName string     `db:"service_name" example:"Netflix"`
	Price       int        `db:"price" example:"322"`
	UserUUID    uuid.UUID  `db:"user_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	StartDate   time.Time  `db:"start_date" example:"03-2024"`
	EndDate     *time.Time `db:"end_date" example:"05-2024"`
}

// SubscriptionSummary represents aggregated subscription data for reporting
// @Description Summary model containing aggregated subscription information for users
type SubscriptionSummary struct {
	TotalPrice  int    `json:"total_price" db:"sum" example:"1337"`
	UserID      string `json:"user_id" db:"user_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
	ServiceName string `json:"service_name" db:"service_name" example:"Netflix"`
}
