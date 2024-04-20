package ordercreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

//go:generate mockgen -destination=./../../../mocks/ordercreatorymock/order_creator_mock.go -source=./order_creator.go -package=ordercreatorymock
type OrderCreator interface {
	// Create(ctx context.Context, order *entities.Order) error
	CreateOrderAndEnqueuePayment(ctx context.Context, order *entities.Order) error
}
