package mapper

import (
	"emtesttask/entity"
	"emtesttask/model"
)

type SubscriptionMapper interface {
	MapEntityToDTO(entity *entity.Subscription) (*model.SubscriptionDTO, error)
	MapDTOToEntity(dto *model.SubscriptionDTO) (*entity.Subscription, error)
}

type subscriptionMapper struct{}

func NewSubscriptionMapper() SubscriptionMapper {
	return &subscriptionMapper{}
}

func (m *subscriptionMapper) MapEntityToDTO(entity *entity.Subscription) (*model.SubscriptionDTO, error) {
	return &model.SubscriptionDTO{
		ServiceName: entity.ServiceName,
		Price:       entity.Price,
		UserId:      entity.UserId,
		StartDate:   entity.StartDate,
		EndDate:     entity.EndDate,
	}, nil
}

func (m *subscriptionMapper) MapDTOToEntity(dto *model.SubscriptionDTO) (*entity.Subscription, error) {
	return &entity.Subscription{
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		UserId:      dto.UserId,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}, nil
}
