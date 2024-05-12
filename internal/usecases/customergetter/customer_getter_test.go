package customergetter_test

import (
	"context"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/customerrepositorymock"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customergetter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ConsumerGetter(t *testing.T) {
	tests := []struct {
		name          string
		input         *entities.Customer
		output        *entities.Customer
		expectedError error
		executeMock   func(input, output *entities.Customer, m *customerrepositorymock.MockCustomerRepository)
	}{
		{
			name:          "Should get customer with success",
			expectedError: nil,
			input: &entities.Customer{
				Document: "52411797044",
			},
			output: &entities.Customer{
				ID:       "fake-id",
				Name:     "John Doe",
				Email:    "mock@mock.com",
				Document: "52411797044",
			},
			executeMock: func(input, output *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(output, nil)
			},
		},
		{
			name:          "Should return not found",
			expectedError: customergetter.ErrCustomerNotFound,
			input: &entities.Customer{
				Document: "52411797044",
			},
			output: nil,
			executeMock: func(input, output *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, nil)
			},
		},
		{
			name:          "Should not get invalid document",
			expectedError: entities.ErrCustomerInvalidDocument,
			input: &entities.Customer{
				Document: "XXXXXX",
			},
			executeMock: func(input, output *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {},
		},
		{
			name:          "Should expected an error to get customer",
			expectedError: assert.AnError,
			input: &entities.Customer{
				Document: "52411797044",
			},
			executeMock: func(input, output *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {
				m.EXPECT().GetByDocument(gomock.Any(), input.Document).Return(nil, assert.AnError)
			},
		},
		{
			name:          "Should expect bad request when document is not passed to get customer",
			expectedError: entities.ErrCustomerEmptyDocument,
			input: &entities.Customer{
				Document: "",
			},
			executeMock: func(input, output *entities.Customer, m *customerrepositorymock.MockCustomerRepository) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepositoryMock := customerrepositorymock.NewMockCustomerRepository(ctrl)

			tt.executeMock(tt.input, tt.output, customerRepositoryMock)

			customerGetter := customergetter.NewCustomerGetter(customerRepositoryMock)

			customer, err := customerGetter.Get(context.Background(), tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.output, customer)
		})
	}
}
