package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

//go:generate mockgen -destination=./../../mocks/customerrepositorymock/customer_repository_mock.go -source=./customer.go -package=customerrepositorymock
type CustomerRepository interface {
	Save(ctx context.Context, user *entities.Customer) error
	GetByDocument(ctx context.Context, document string) (*entities.Customer, error)
	GetByDocumentAndEmail(ctx context.Context, document string, email string) (*entities.Customer, error)
}
