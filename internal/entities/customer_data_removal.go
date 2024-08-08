package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCustomerDataRemoval = errors.New("invalid customer data removal")
)

type CustomerDataRemovalRequest struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	Phonenumber      string    `json:"phonenumber"`
	Document         string    `json:"document"`
	CreatedDate      time.Time `json:"created_date_db"`
	LastModifiedDate time.Time `json:"last_modified_date_db"`
}

func NewCustomerDataRemovalRequest(name string, address string, phonenumber string, document string) *CustomerDataRemovalRequest {
	return &CustomerDataRemovalRequest{
		ID:               uuid.NewString(),
		Name:             name,
		Address:          address,
		Phonenumber:      phonenumber,
		Document:         document,
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}
}

func (u *CustomerDataRemovalRequest) Validate() error {
	if u.Name == "" {
		return ErrCustomerDataRemoval
	}

	if u.Address == "" {
		return ErrCustomerDataRemoval
	}

	if u.Phonenumber == "" {
		return ErrCustomerDataRemoval
	}

	return nil
}
