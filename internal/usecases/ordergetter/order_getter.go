package ordergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

type OrderGetter struct {
	orderRepository   interfaces.OrderRepository
	productRepository interfaces.ProductRepository
}

func NewOrderGetter(orderRepository interfaces.OrderRepository, productRepository interfaces.ProductRepository) *OrderGetter {
	return &OrderGetter{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (p *OrderGetter) GetAll(ctx context.Context, sortBy string) ([]entities.Order, error) {
	if sortBy == "status" {
		return p.orderRepository.GetAllSortedByStatus(ctx)
	}

	list, err := p.orderRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}
