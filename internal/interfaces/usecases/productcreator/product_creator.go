package productcreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

type ProductCreator interface {
	Add(ctx context.Context, product *entities.Product) error
}
