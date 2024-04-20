package customergetterhandler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	customergetterhandler "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/customergetterandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/customergetterymock"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customergetter"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_CustomerGetter_Handler(t *testing.T) {

	tests := []struct {
		name               string
		inputBody          string
		expectedStatusCode int
		executeMock        func(m *customergetterymock.MockCustomerGetter)
	}{
		{
			name:               "Should customer get with success",
			inputBody:          `1234567`,
			expectedStatusCode: http.StatusOK,
			executeMock: func(m *customergetterymock.MockCustomerGetter) {
				m.EXPECT().Get(gomock.Any(), &entities.Customer{Document: "1234567"}).Return(&entities.Customer{ID: "fake-id"}, nil)
			},
		},
		{
			name:               "Should return not found",
			inputBody:          `1234567`,
			expectedStatusCode: http.StatusNotFound,
			executeMock: func(m *customergetterymock.MockCustomerGetter) {
				m.EXPECT().Get(gomock.Any(), &entities.Customer{Document: "1234567"}).Return(nil, customergetter.ErrCustomerNotFound)
			},
		},
		{
			name:               "Should invalid document",
			inputBody:          `1234567`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *customergetterymock.MockCustomerGetter) {
				m.EXPECT().Get(gomock.Any(), &entities.Customer{Document: "1234567"}).Return(nil, entities.ErrCustomerInvalidDocument)
			},
		},
		{
			name:               "Should return an general error",
			inputBody:          `1234567`,
			expectedStatusCode: http.StatusInternalServerError,
			executeMock: func(m *customergetterymock.MockCustomerGetter) {
				m.EXPECT().Get(gomock.Any(), &entities.Customer{Document: "1234567"}).Return(nil, assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			customerGetterUseCaseMock := customergetterymock.NewMockCustomerGetter(ctrl)

			tt.executeMock(customerGetterUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/customers"
			customerGetterHandler := customergetterhandler.NewCustomerGetterHandler(customerGetterUseCaseMock)

			appFake.GET(endpoint, customerGetterHandler.Handle)
			req, _ := http.NewRequest("GET", fmt.Sprintf("%s?document=%s", endpoint, tt.inputBody), nil)
			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
