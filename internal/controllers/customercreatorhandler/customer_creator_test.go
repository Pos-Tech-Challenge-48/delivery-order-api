package customercreatorhandler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/customercreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/customercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/customercreatorymock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CustomerCreator_Handler(t *testing.T) {

	tests := []struct {
		name               string
		inputBody          string
		expectedStatusCode int
		executeMock        func(m *customercreatorymock.MockCustomerCreator)
	}{
		{
			name:               "Should customer created with success",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusCreated,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(nil)
			},
		},
		{
			name:               "Should return empty name",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(entities.ErrCustomerEmptyName)
			},
		},
		{
			name:               "Should return empty email",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(entities.ErrCustomerEmptyEmail)
			},
		},
		{
			name:               "Should return already exists customer",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(customercreator.ErrAlreadyExistsCustomer)
			},
		},
		{
			name:               "Should return invalid document",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(entities.ErrCustomerInvalidDocument)
			},
		},
		{
			name:               "Should return invalid email",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(entities.ErrCustomerInvalidEmail)
			},
		},
		{
			name:               "Should return an general error",
			inputBody:          `{"name": "John Doe","email": "mock@mock.com","document":"XXXXX"}`,
			expectedStatusCode: http.StatusInternalServerError,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
				customer := &entities.Customer{Name: "John Doe", Email: "mock@mock.com", Document: "XXXXX"}
				m.EXPECT().Create(gomock.Any(), customer).Return(assert.AnError)
			},
		},
		{
			name:               "Should return an error to bind",
			inputBody:          `{`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customercreatorymock.MockCustomerCreator) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			customerCreatorUseCaseMock := customercreatorymock.NewMockCustomerCreator(ctrl)

			tt.executeMock(customerCreatorUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/customers"
			customerCreatorHandler := customercreatorhandler.NewCustomerCreatorHandler(customerCreatorUseCaseMock)

			appFake.POST(endpoint, customerCreatorHandler.Handle)
			req, _ := http.NewRequest("POST", endpoint, strings.NewReader(tt.inputBody))
			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
