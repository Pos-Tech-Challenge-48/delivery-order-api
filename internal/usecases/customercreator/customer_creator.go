package customercreator

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

var (
	ErrAlreadyExistsCustomer = errors.New("already exists customer with this document")
)

type CustomerCreator struct {
	customerRepository interfaces.CustomerRepository
}

func NewCustomerCreator(customerRepository interfaces.CustomerRepository) *CustomerCreator {
	return &CustomerCreator{
		customerRepository: customerRepository,
	}
}

func (uc *CustomerCreator) Create(ctx context.Context, customerInput *entities.Customer) error {

	err := customerInput.Validate()
	if err != nil {
		return err
	}

	existsCustomer, err := uc.customerRepository.GetByDocument(ctx, customerInput.Document)
	if err != nil {
		return err
	}

	if existsCustomer != nil {
		return ErrAlreadyExistsCustomer
	}

	customer := entities.NewCustomer(customerInput.Name, customerInput.Email, customerInput.Document)

	return uc.customerRepository.Save(ctx, customer)
}
