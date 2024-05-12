package customercreator

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
)

var (
	ErrAlreadyExistsCustomer = errors.New("already exists customer with this document")
)

//go:generate mockgen -destination=./../../../mocks/customercreatorymock/customer_creator_mock.go -source=./customer_creator.go -package=customercreatorymock
type CustomerCreator interface {
	Create(ctx context.Context, user *entities.Customer) error
}
