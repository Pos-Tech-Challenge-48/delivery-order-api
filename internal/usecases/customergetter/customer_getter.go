package customergetter

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

var (
	ErrCustomerNotFound = errors.New("customer not find with this document")
)

type CustomerGetter struct {
	customerRepository interfaces.CustomerRepository
}

func NewCustomerGetter(customerRepository interfaces.CustomerRepository) *CustomerGetter {
	return &CustomerGetter{
		customerRepository: customerRepository,
	}
}

func (uc *CustomerGetter) Get(ctx context.Context, customerInput *entities.Customer) (*entities.Customer, error) {
	err := customerInput.ValidateDocument()
	if err != nil {
		return nil, err
	}

	customer, err := uc.customerRepository.GetByDocument(ctx, customerInput.Document)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, ErrCustomerNotFound
	}

	return customer, nil
}
