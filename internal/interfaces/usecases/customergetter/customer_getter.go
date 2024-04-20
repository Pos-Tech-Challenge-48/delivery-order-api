package customergetter

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

//go:generate mockgen -destination=./../../../mocks/customergetterymock/customer_getter_mock.go -source=./customer_getter.go -package=customergetterymock
type CustomerGetter interface {
	Get(ctx context.Context, customerInput *entities.Customer) (*entities.Customer, error)
}
