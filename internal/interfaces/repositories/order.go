package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

//go:generate mockgen -destination=./../../mocks/orderrepositorymock/order_repository_mock.go -source=./order.go -package=orderrepositorymock
type OrderRepository interface {
	Save(ctx context.Context, order *entities.Order) error
	GetAll(ctx context.Context) ([]entities.Order, error)
	GetAllSortedByStatus(ctx context.Context) ([]entities.Order, error)
	GetByID(ctx context.Context, orderID string) (*entities.Order, error)
	Update(ctx context.Context, order *entities.Order) error
}

type OrderQueueRepository interface {
	SendPendingPaymentOrderMessageToQueue(ctx context.Context, order *entities.Order) error
}
