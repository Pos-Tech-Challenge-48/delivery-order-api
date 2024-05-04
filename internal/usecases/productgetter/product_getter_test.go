package productgetter_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productgetter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ProductGetter(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		repository    []entities.Product
		output        []entities.Product
		expectedError error
		executeMock   func(input string, repository []entities.Product, m *productrepositorymock.MockProductRepository)
	}{
		{
			name:          "Should create new product with success",
			expectedError: nil,
			input:         "Uuid-Bebida",
			repository: []entities.Product{
				{
					Name:       "Coca",
					CategoryID: "Uuid-Bebida",
				},
				{
					Name:       "Lanche",
					CategoryID: "Uuid-Comida",
				},
			},
			output: []entities.Product{
				{
					Name:       "Coca",
					CategoryID: "Uuid-Bebida",
				},
			},
			executeMock: func(input string, repository []entities.Product, m *productrepositorymock.MockProductRepository) {
				expectedOutput := []entities.Product{}
				if input == "" {
					expectedOutput = repository
				} else {
					for _, product := range repository {
						if product.CategoryID == input {
							expectedOutput = append(expectedOutput, product)
						}
					}
				}
				m.EXPECT().GetAll(gomock.Any(), input).Return(expectedOutput, nil)
			},
		},

		{
			name:          "Empty List",
			expectedError: nil,
			input:         "Uuid-Bebida",
			repository:    []entities.Product{},
			output:        nil,
			executeMock: func(input string, repository []entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().GetAll(gomock.Any(), input).Return(nil, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := productrepositorymock.NewMockProductRepository(gomock.NewController(t))
			tt.executeMock(tt.input, tt.repository, m)

			ProductGetter := productgetter.NewProductGetter(m)

			productList, err := ProductGetter.GetAll(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, productList)
		})
	}
}
