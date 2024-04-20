package ordercreator_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/orderrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/ordercreator"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderCreator(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Order
		expectedError error
		executeMock   func(input *entities.Order, m *orderrepositorymock.MockOrderRepository,
			pm *productrepositorymock.MockProductRepository)
	}{
		{
			name:          "Should create order with success",
			expectedError: nil,
			input: &entities.Order{
				CustomerID: "uuid",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "uuid",
						Name: "Coca",
					},
					{
						ID:   "uuid2",
						Name: "X-burger",
					},
				},
			},
			executeMock: func(input *entities.Order, m *orderrepositorymock.MockOrderRepository, pm *productrepositorymock.MockProductRepository) {
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
				pm.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&entities.Product{}, nil).MaxTimes(2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			OrderRepositoryMock := orderrepositorymock.NewMockOrderRepository(ctrl)

			productRepositoryMock := productrepositorymock.NewMockProductRepository(ctrl)
			NewMockOrderQueueRepository := orderrepositorymock.NewMockOrderQueueRepository(ctrl)

			tt.executeMock(tt.input, OrderRepositoryMock, productRepositoryMock)

			OrderCreator := ordercreator.NewOrderCreator(OrderRepositoryMock, productRepositoryMock, NewMockOrderQueueRepository)

			err := OrderCreator.CreateOrderAndEnqueuePayment(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
