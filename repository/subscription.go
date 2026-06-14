package repository

import (
	"context"
	"emtesttask/entity"
	"emtesttask/model"

	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetSubscriptionByID(ctx context.Context, id uint64) (*entity.Subscription, error)
	AddSubscription(ctx context.Context, newSub *entity.Subscription) (*entity.Subscription, error)
	GetSubscriptionsPaged(ctx context.Context, pagedRequest *model.PagedRequest) ([]entity.Subscription, error)
}

type subscriptionRepository struct {
	db *gorm.DB
}

// Провайдер для Wire
func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) GetSubscriptionByID(ctx context.Context, id uint64) (*entity.Subscription, error) {
	var subscription entity.Subscription
	result := r.db.WithContext(ctx).First(&subscription, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subscription, nil
}

func (r *subscriptionRepository) AddSubscription(ctx context.Context, newSub *entity.Subscription) (*entity.Subscription, error) {
	err := gorm.G[entity.Subscription](r.db).Create(ctx, newSub)
	if err != nil {
		return nil, err
	}
	return newSub, nil
}

func (r *subscriptionRepository) GetSubscriptionsPaged(ctx context.Context, pagedRequest *model.PagedRequest) ([]entity.Subscription, error) {
	result, err := gorm.G[entity.Subscription](r.db).Limit(
		pagedRequest.PageSize,
	).Offset(
		(pagedRequest.PageNumber - 1) * pagedRequest.PageSize,
	).Find(ctx)
	if err != nil {
		return nil, err
	}
	return result, err
}
