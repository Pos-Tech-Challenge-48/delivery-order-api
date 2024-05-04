package productdelete_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productdelete"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ProductDelete(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Product
		expectedError error
		executeMock   func(input *entities.Product, m *productrepositorymock.MockProductRepository)
	}{
		{
			name:          "Should delete Product",
			expectedError: nil,
			input: &entities.Product{
				ID:         "fake-product-id ",
				Name:       "Coca",
				Image:      "fake-image",
				CategoryID: "Uuid",
			},
			executeMock: func(input *entities.Product, m *productrepositorymock.MockProductRepository) {
				m.EXPECT().DeleteImage(gomock.Any(), input.ID).Return(nil)
				m.EXPECT().Delete(gomock.Any(), input.ID).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := productrepositorymock.NewMockProductRepository(gomock.NewController(t))
			tt.executeMock(tt.input, m)

			ProductDeleteAgent := productdelete.NewProductDelete(m)

			err := ProductDeleteAgent.Delete(context.Background(), tt.input.ID)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
