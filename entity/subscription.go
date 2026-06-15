package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	ServiceName string     `gorm:"column:service_name;size:1024"`
	UserId      uuid.UUID  `gorm:"column:user_id"`
	Price       int32      `gorm:"column:price"`
	StartDate   time.Time  `gorm:"column:start_date"`
	EndDate     *time.Time `gorm:"column:end_date"`
}

type SubscriptionSum struct {
	ServiceName string `gorm:"column:sn"`
	Sum         uint64 `gorm:"column:sum_price"`
}
