package productcreator_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productcreator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ProductCreator(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Product
		expectedError error
		executeMock   func(input *entities.Product, m *productrepositorymock.MockProductRepository)
	}{
		{
			name:          "Should create new product with success",
			expectedError: nil,
			input: &entities.Product{
				Name:       "Coca",
				CategoryID: "Uuid",
			},
			executeMock: func(input *entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().GetCategoryID(gomock.Any(), input.CategoryID).Return(input.CategoryID, nil)
				m.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := productrepositorymock.NewMockProductRepository(gomock.NewController(t))
			tt.executeMock(tt.input, m)

			ProductCreator := productcreator.NewProductCreator(m)

			err := ProductCreator.Add(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
