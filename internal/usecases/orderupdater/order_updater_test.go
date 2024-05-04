package orderupdater_test

import (
	"context"
	"testing"
	"time"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/orderrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/orderupdater"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderUpdater(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Order
		expectedError error
		executeMock   func(input *entities.Order, m *orderrepositorymock.MockOrderRepository, q *orderrepositorymock.MockOrderQueueRepository)
	}{
		{
			name:          "Should successfully update an order",
			expectedError: nil,
			input: &entities.Order{
				ID:         "fake-id",
				CustomerID: "fake-customer-id",
				Status:     "Recebido",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "fake-procuct-id",
						Name: "fake-product-name",
					},
				},
				CreatedDate:      time.Now(),
				LastModifiedDate: time.Now(),
			},
			executeMock: func(input *entities.Order, m *orderrepositorymock.MockOrderRepository, q *orderrepositorymock.MockOrderQueueRepository) {
				m.EXPECT().GetByID(gomock.Any(), input.ID).Return(input, nil)
				m.EXPECT().Update(gomock.Any(), input).Return(nil)
			},
		},
		{
			name:          "Order cannot be updated because it does not exist",
			expectedError: entities.ErrOrderNotExist,
			input: &entities.Order{
				ID: "NON EXISTENT ORDER",
			},
			executeMock: func(input *entities.Order, m *orderrepositorymock.MockOrderRepository, q *orderrepositorymock.MockOrderQueueRepository) {
				m.EXPECT().GetByID(gomock.Any(), input.ID).Return(nil, nil)
			},
		},
		{
			name:          "Order cannot be updated because it's finished",
			expectedError: entities.ErrInvalidOrderStatus,
			input: &entities.Order{
				ID:         "fake-id",
				CustomerID: "fake-customer-id",
				Status:     "Finalizado",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "fake-procuct-id",
						Name: "fake-product-name",
					},
				},
				CreatedDate:      time.Now(),
				LastModifiedDate: time.Now(),
			},
			executeMock: func(input *entities.Order, m *orderrepositorymock.MockOrderRepository, q *orderrepositorymock.MockOrderQueueRepository) {
				m.EXPECT().GetByID(gomock.Any(), input.ID).Return(input, nil)
			},
		},
		{
			name:          "Order is being updated and should be enqueued",
			expectedError: nil,
			input: &entities.Order{
				ID:         "fake-id",
				CustomerID: "fake-customer-id",
				Status:     "Pago",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "fake-procuct-id",
						Name: "fake-product-name",
					},
				},
				CreatedDate:      time.Now(),
				LastModifiedDate: time.Now(),
			},
			executeMock: func(input *entities.Order, m *orderrepositorymock.MockOrderRepository, q *orderrepositorymock.MockOrderQueueRepository) {
				m.EXPECT().GetByID(gomock.Any(), input.ID).Return(input, nil)
				m.EXPECT().Update(gomock.Any(), input).Return(nil)
				q.EXPECT().SendOrderToProductionQueue(gomock.Any(), input).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepositoryMock := orderrepositorymock.NewMockOrderRepository(ctrl)
			orderEnqueuerMock := orderrepositorymock.NewMockOrderQueueRepository(ctrl)

			tt.executeMock(tt.input, orderRepositoryMock, orderEnqueuerMock)

			orderUpdater := orderupdater.NewOrderUpdater(orderRepositoryMock, orderEnqueuerMock)

			err := orderUpdater.Update(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
		})
	}

}
