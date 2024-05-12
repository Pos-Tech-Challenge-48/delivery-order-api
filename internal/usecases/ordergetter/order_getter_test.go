package ordergetter_test

import (
	"context"
	"testing"
	"time"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/orderrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/ordergetter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// mock tem que ser do repository de order
var createdDate = time.Now()

func Test_OrderGetter(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		repository    []entities.Order
		output        []entities.Order
		expectedError error
		executeMock   func(input string, repository []entities.Order, output []entities.Order, m *orderrepositorymock.MockOrderRepository)
	}{
		{
			name:          "Should return a OrderList",
			expectedError: nil,
			input:         "",
			repository: []entities.Order{
				{
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
			},
			output: []entities.Order{
				{
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
			},
			executeMock: func(input string, repository []entities.Order, output []entities.Order, m *orderrepositorymock.MockOrderRepository) {
				m.EXPECT().GetAll(gomock.Any()).Return(output, nil)
			},
		},
		{
			name:          "Should return one order",
			expectedError: nil,
			input:         "status",
			repository: []entities.Order{
				{
					ID:         "fake-id",
					CustomerID: "fake-customer-id",
					Status:     "Recebido",
					OrderProduct: []entities.OrderProduct{
						{
							ID:   "fake-procuct-id",
							Name: "fake-product-name",
						},
					},
					CreatedDate:      createdDate,
					LastModifiedDate: createdDate,
				},
				{
					ID:         "fake-id-ready",
					CustomerID: "fake-customer-id",
					Status:     "Pronto",
					OrderProduct: []entities.OrderProduct{
						{
							ID:   "fake-procuct-id",
							Name: "fake-product-name",
						},
					},
					CreatedDate:      createdDate,
					LastModifiedDate: createdDate,
				},
			},
			output: []entities.Order{
				{
					ID:         "fake-id-ready",
					CustomerID: "fake-customer-id",
					Status:     "Pronto",
					OrderProduct: []entities.OrderProduct{
						{
							ID:   "fake-procuct-id",
							Name: "fake-product-name",
						},
					},
					CreatedDate:      createdDate,
					LastModifiedDate: createdDate,
				},
				{
					ID:         "fake-id",
					CustomerID: "fake-customer-id",
					Status:     "Recebido",
					OrderProduct: []entities.OrderProduct{
						{
							ID:   "fake-procuct-id",
							Name: "fake-product-name",
						},
					},
					CreatedDate:      createdDate,
					LastModifiedDate: createdDate,
				},
			},
			executeMock: func(input string, repository []entities.Order, output []entities.Order, m *orderrepositorymock.MockOrderRepository) {
				m.EXPECT().GetAllSortedByStatus(gomock.Any()).Return([]entities.Order{repository[1], repository[0]}, nil)
			},
		},
		{
			name:          "Empty list",
			expectedError: nil,
			input:         "",
			repository:    []entities.Order{},
			output:        nil,
			executeMock: func(input string, repository []entities.Order, output []entities.Order, m *orderrepositorymock.MockOrderRepository) {
				m.EXPECT().GetAll(gomock.Any()).Return(output, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepositoryMock := orderrepositorymock.NewMockOrderRepository(ctrl)

			productRepositoryMock := productrepositorymock.NewMockProductRepository(ctrl)

			tt.executeMock(tt.input, tt.repository, tt.output, orderRepositoryMock)

			orderGetter := ordergetter.NewOrderGetter(orderRepositoryMock, productRepositoryMock)

			orderList, err := orderGetter.GetAll(context.Background(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, orderList)
		})
	}

}
