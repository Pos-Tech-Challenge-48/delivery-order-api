package productdelete

import (
	"context"
	"fmt"

	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

type ProductDelete struct {
	productRepository interfaces.ProductRepository
}

func NewProductDelete(productRepository interfaces.ProductRepository) *ProductDelete {
	return &ProductDelete{
		productRepository: productRepository,
	}
}

func (p *ProductDelete) Delete(ctx context.Context, producID string) error {

	err := p.productRepository.DeleteImage(ctx, producID)
	if err != nil {
		return fmt.Errorf("error on delete image product: %w", err)
	}

	err = p.productRepository.Delete(ctx, producID)
	if err != nil {
		return fmt.Errorf("error to submit delete product: %w", err)
	}

	return nil
}
