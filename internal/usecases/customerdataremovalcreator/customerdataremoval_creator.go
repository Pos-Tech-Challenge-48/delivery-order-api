package customerdataremovalcreator

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

var (
	ErrAlreadyExistsCustomer = errors.New("already exists customer with this document")
)

type CustomerDataRemovalRequestCreator struct {
	cdrRepository interfaces.CustomerDataRemovalRepository
}

func NewCustomerDataRemovalCreator(cdrRepository interfaces.CustomerDataRemovalRepository) *CustomerDataRemovalRequestCreator {
	return &CustomerDataRemovalRequestCreator{
		cdrRepository: cdrRepository,
	}
}

func (uc *CustomerDataRemovalRequestCreator) Create(ctx context.Context, data *entities.CustomerDataRemovalRequest) error {

	err := data.Validate()
	if err != nil {
		return err
	}

	customerDataRemovalRequest := entities.NewCustomerDataRemovalRequest(data.Name, data.Address, data.Phonenumber, data.Document)

	err = uc.cdrRepository.Save(ctx, customerDataRemovalRequest)
	if err != nil {
		return err
	}

	return nil
}
