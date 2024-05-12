package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

//go:generate mockgen -destination=./../../mocks/productrepositorymock/product_repository_mock.go -source=./product.go -package=productrepositorymock
type ProductRepository interface {
	Add(ctx context.Context, product *entities.Product) error
	Update(ctx context.Context, product *entities.Product) error
	UpdateImage(ctx context.Context, productID string, image string) error
	Delete(ctx context.Context, productID string) error
	GetAll(ctx context.Context, params string) ([]entities.Product, error)
	GetByID(ctx context.Context, ID string) (*entities.Product, error)
	GetCategoryID(ctx context.Context, categoryName string) (categoryID string, err error)
	DeleteImage(ctx context.Context, productID string) (err error)
}
