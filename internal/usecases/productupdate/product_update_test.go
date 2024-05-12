package productupdate_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productupdate"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ProductUpdater(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Product
		repository    []entities.Product
		output        []entities.Product
		expectedError error
		executeMock   func(input *entities.Product, repository []entities.Product, m *productrepositorymock.MockProductRepository)
	}{
		{
			name:          "Should create new product with success",
			expectedError: nil,
			input: &entities.Product{
				ID:         "product-id",
				Name:       "Coca",
				CategoryID: "Uuid-Bebida",
				Image:      "fake-image-url",
			},
			executeMock: func(input *entities.Product, repository []entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().GetCategoryID(gomock.Any(), input.CategoryID).Return(input.CategoryID, nil)
				m.EXPECT().Update(gomock.Any(), input).Return(nil)
				m.EXPECT().UpdateImage(gomock.Any(), input.ID, input.Image).Return(nil)
			},
		},
		{
			name:          "Error fetching category",
			expectedError: errors.New(""),
			input:         &entities.Product{},
			executeMock: func(input *entities.Product, repository []entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().GetCategoryID(gomock.Any(), input.CategoryID).Return("", errors.New(""))
			},
		},
		{
			name:          "Error updating",
			expectedError: errors.New(""),
			input:         &entities.Product{},
			executeMock: func(input *entities.Product, repository []entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().GetCategoryID(gomock.Any(), input.CategoryID).Return(input.CategoryID, nil)
				m.EXPECT().Update(gomock.Any(), input).Return(errors.New(""))
			},
		},
		{
			name:          "Error updating image",
			expectedError: errors.New(""),
			input:         &entities.Product{},
			executeMock: func(input *entities.Product, repository []entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().GetCategoryID(gomock.Any(), input.CategoryID).Return(input.CategoryID, nil)
				m.EXPECT().Update(gomock.Any(), input).Return(nil)
				m.EXPECT().UpdateImage(gomock.Any(), input.ID, input.Image).Return(errors.New(""))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := productrepositorymock.NewMockProductRepository(gomock.NewController(t))
			tt.executeMock(tt.input, tt.repository, m)
			ProductUpdator := productupdate.NewProductUpdate(m)

			err := ProductUpdator.Update(context.Background(), tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
