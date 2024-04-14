package customercreator_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/customerrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customercreator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ConsumerCreator(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Customer
		expectedError error
		executeMock   func(input *entities.Customer, m *customerrepositorymock.MockCustomerRepository)
	}{
		{
			name:          "Should create customerr with success",
			expectedError: nil,
			input: &entities.Customer{
				Name:     "John Doe",
				Email:    "mock@mock.com",
				Document: "52411797044",
			},
			executeMock: func(input *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, nil)
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:          "Should not create invalid customer",
			expectedError: entities.ErrCustomerEmptyEmail,
			input: &entities.Customer{
				Name:     "John Doe",
				Document: "52411797044",
			},
			executeMock: func(input *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {},
		},
		{
			name:          "Should expected an error to get customer",
			expectedError: assert.AnError,
			input: &entities.Customer{
				Name:     "John Doe",
				Email:    "mock@mock.com",
				Document: "52411797044",
			},
			executeMock: func(input *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, assert.AnError)
			},
		},
		{
			name:          "Should return an error to create customer",
			expectedError: assert.AnError,
			input: &entities.Customer{
				Name:     "John Doe",
				Email:    "mock@mock.com",
				Document: "52411797044",
			},
			executeMock: func(input *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, nil)
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
		},
		{
			name:          "Should return already exists customer with this document",
			expectedError: customercreator.ErrAlreadyExistsCustomer,
			input: &entities.Customer{
				Name:     "John Doe",
				Email:    "mock@mock.com",
				Document: "52411797044",
			},
			executeMock: func(input *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(input, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepositoryMock := customerrepositorymock.NewMockCustomerRepository(ctrl)

			tt.executeMock(tt.input, customerRepositoryMock)

			customerCreator := customercreator.NewCustomerCreator(customerRepositoryMock)

			err := customerCreator.Create(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
