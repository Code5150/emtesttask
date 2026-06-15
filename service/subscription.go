package service

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/samber/lo"

	"emtesttask/entity"
	"emtesttask/mapper"
	"emtesttask/model"
	"emtesttask/repository"
)

type SubscriptionService interface {
	GetSubscriptionByID(ctx context.Context, id uint64) (*model.SubscriptionDTO, error)
	GetSubscriptionsPaged(ctx context.Context, pagedRequest *model.PagedRequest) ([]model.SubscriptionDTO, error)
	GetSubscriptionsSum(ctx context.Context, subscriptionFilter *model.SubscriptionSumRequest) (*model.SubscriptionSumResponse, error)
	AddSubscription(ctx context.Context, newSub *model.SubscriptionDTO) (*model.SubscriptionDTO, error)
	UpdateSubscription(ctx context.Context, id uint64, newSub *model.SubscriptionDTO) (*model.SubscriptionDTO, error)
	DeleteSubscription(ctx context.Context, id uint64) error
}

type subscriptionService struct {
	repo   repository.SubscriptionRepository
	mapper mapper.SubscriptionMapper
}

func NewSubscriptionService(repo repository.SubscriptionRepository, mapper mapper.SubscriptionMapper) SubscriptionService {
	return &subscriptionService{repo: repo, mapper: mapper}
}

func (s *subscriptionService) GetSubscriptionByID(ctx context.Context, id uint64) (*model.SubscriptionDTO, error) {
	subscription, err := s.repo.GetSubscriptionByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Subscription not found")
		}
		return nil, err
	}
	dao, err := s.mapper.MapEntityToDTO(subscription)
	if err != nil {
		log.Print("Failed to load entity", err)
		return nil, err
	}
	return dao, nil
}

func (s *subscriptionService) GetSubscriptionsPaged(ctx context.Context, pagedRequest *model.PagedRequest) ([]model.SubscriptionDTO, error) {
	subscriptions, err := s.repo.GetSubscriptionsPaged(ctx, pagedRequest)
	if err != nil {
		log.Print("Failed to load entity", err)
		return nil, err
	}
	result, err := lo.MapErr(subscriptions, func(sub entity.Subscription, _ int) (model.SubscriptionDTO, error) {
		result, err := s.mapper.MapEntityToDTO(&sub)
		if err != nil {
			log.Print("Failed to load entity", err)
			return model.SubscriptionDTO{}, err
		}
		return *result, nil
	})
	if err != nil {
		log.Print("Failed to load entity", err)
		return nil, err
	}
	return result, nil
}

func (s *subscriptionService) GetSubscriptionsSum(ctx context.Context, subscriptionFilter *model.SubscriptionSumRequest) (*model.SubscriptionSumResponse, error) {
	log.Print(subscriptionFilter)
	return s.repo.GetSubscriptionsSum(ctx, subscriptionFilter)
}

func (s *subscriptionService) AddSubscription(ctx context.Context, newSub *model.SubscriptionDTO) (*model.SubscriptionDTO, error) {
	log.Println(newSub)
	entity, err := s.mapper.MapDTOToEntity(newSub)
	if err != nil {
		log.Print("Failed to create entity", err)
		return nil, err
	}
	log.Println(entity)
	result, err := s.repo.AddSubscription(ctx, entity)
	if err != nil {
		log.Print("Failed to insert entity", err)
		return nil, err
	}
	return s.mapper.MapEntityToDTO(result)
}

func (s *subscriptionService) UpdateSubscription(ctx context.Context, id uint64, newSub *model.SubscriptionDTO) (*model.SubscriptionDTO, error) {
	log.Println(newSub)
	entity, err := s.mapper.MapDTOToEntity(newSub)
	if err != nil {
		log.Print("Failed to create entity", err)
		return nil, err
	}
	log.Println(entity)
	result, err := s.repo.UpdateSubscription(ctx, id, entity)
	if err != nil {
		log.Print("Failed to insert entity", err)
		return nil, err
	}
	return s.mapper.MapEntityToDTO(result)
}

func (s *subscriptionService) DeleteSubscription(ctx context.Context, id uint64) error {
	log.Println("Deleting id ", id)
	err := s.repo.DeleteSubscription(ctx, id)
	if err != nil {
		log.Print("Failed to delete entity", err)
		return err
	}
	return nil
}
