package ordercreatorhandler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/ordercreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/ordercreatorymock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderCreator_Handler(t *testing.T) {

	fakeOrder := &entities.Order{
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
		}}

	tests := []struct {
		name               string
		inputBody          string
		expectedStatusCode int
		executeMock        func(m *ordercreatorymock.MockOrderCreator)
	}{
		{
			name:               "Should Order created with success",
			inputBody:          `{"customer_id": "uuid","products": [{"id":"uuid","name":"Coca"},{"id":"uuid2","name":"X-burger"}]}`,
			expectedStatusCode: http.StatusCreated,
			executeMock: func(m *ordercreatorymock.MockOrderCreator) {
				m.EXPECT().CreateOrderAndEnqueuePayment(gomock.Any(), fakeOrder).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			OrderCreatorUseCaseMock := ordercreatorymock.NewMockOrderCreator(ctrl)

			tt.executeMock(OrderCreatorUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/orders"
			OrderCreatorHandler := ordercreatorhandler.NewOrderCreatorHandler(OrderCreatorUseCaseMock)

			appFake.POST(endpoint, OrderCreatorHandler.Handle)
			req, _ := http.NewRequest("POST", endpoint, strings.NewReader(tt.inputBody))
			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
