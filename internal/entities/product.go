package entities

import (
	"errors"
	"time"
)

type Product struct {
	ID               string    `json:"id"`
	CategoryID       string    `json:"category"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	Image            string    `json:"image"`
	CreatedDate      time.Time `json:"created_date_db"`
	LastModifiedDate time.Time `json:"last_modified_date_db"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("name: can't be empty")
	}

	if p.CategoryID == "" {
		return errors.New("category: can't be empty")
	}

	if p.Description == "" {
		return errors.New("description: can't be empty")
	}

	if p.Price <= 0 {
		return errors.New("price: invalid")
	}

	if p.Image == "" {
		return errors.New("image: can't be empty")
	}

	return nil
}
