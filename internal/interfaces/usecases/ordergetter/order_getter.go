package ordergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

type OrderGetter interface {
	GetAll(ctx context.Context, sortBy string) ([]entities.Order, error)
}
