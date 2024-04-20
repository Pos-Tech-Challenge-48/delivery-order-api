package productgetter

import (
	"context"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

type ProductGetter struct {
	productRepository interfaces.ProductRepository
}

func NewProductGetter(productRepository interfaces.ProductRepository) *ProductGetter {
	return &ProductGetter{
		productRepository: productRepository,
	}
}

func (p *ProductGetter) GetAll(ctx context.Context, category string) ([]entities.Product, error) {

	formatterTitle := cases.Title(language.Portuguese)

	categoryName := formatterTitle.String(strings.ToLower(category))

	list, err := p.productRepository.GetAll(ctx, categoryName)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}
