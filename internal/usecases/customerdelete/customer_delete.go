package customerdelete

import (
	"context"
	"errors"

	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

var ErrCustomerNotFound = errors.New("customer not found")

type CustomerDelete struct {
	customerRepository interfaces.CustomerRepository
}

func NewCustomerDelete(customerRepository interfaces.CustomerRepository) *CustomerDelete {
	return &CustomerDelete{
		customerRepository: customerRepository,
	}
}

func (uc *CustomerDelete) Handle(ctx context.Context, ID string) error {
	customer, err := uc.customerRepository.GetByID(ctx, ID)
	if err != nil {
		return err
	}

	if customer == nil {
		return ErrCustomerNotFound
	}

	err = uc.customerRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
