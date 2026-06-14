package model

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionDTO struct {
	ServiceName string     `json:"service_name"`
	Price       int32      `json:"price"`
	UserId      uuid.UUID  `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitzero"`
}
