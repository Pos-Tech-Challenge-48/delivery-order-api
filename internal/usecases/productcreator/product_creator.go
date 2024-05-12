package productcreator

import (
	"context"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ProductCreator struct {
	productRepository interfaces.ProductRepository
}

func NewProductCreator(productRepository interfaces.ProductRepository) *ProductCreator {
	return &ProductCreator{
		productRepository: productRepository,
	}
}

func (p *ProductCreator) Add(ctx context.Context, data *entities.Product) error {

	formatterTitle := cases.Title(language.Portuguese)

	categoryName := formatterTitle.String(strings.ToLower(data.CategoryID))

	categoryId, err := p.productRepository.GetCategoryID(ctx, categoryName)
	if err != nil {
		return err
	}

	data.CategoryID = categoryId
	err = p.productRepository.Add(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
