package model

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionSumRequest struct {
	UserId      uuid.UUID  `json:"user_id"`
	ServiceName string     `json:"service_name"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitzero"`
}

type SubscriptionSumResponse struct {
	SumPrice uint64 `json:"sum"`
}
