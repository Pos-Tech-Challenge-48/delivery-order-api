package ordergetterhandler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/ordergetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	ordergetterymock "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/ordergettermock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderGetter_Handler(t *testing.T) {

	fakeOrderList := []entities.Order{
		{
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
	}

	tests := []struct {
		name       string
		inputParam string

		expectedStatusCode int
		executeMock        func(inputParam string, m *ordergetterymock.MockOrderGetter)
	}{
		{
			name:               "Should return order from database",
			inputParam:         "status",
			expectedStatusCode: http.StatusOK,
			executeMock: func(inputParam string, m *ordergetterymock.MockOrderGetter) {
				m.EXPECT().GetAll(gomock.Any(), inputParam).Return(fakeOrderList, nil)
			},
		},
		{
			name:               "Should fail correctly",
			inputParam:         "status",
			expectedStatusCode: http.StatusInternalServerError,
			executeMock: func(inputParam string, m *ordergetterymock.MockOrderGetter) {
				m.EXPECT().GetAll(gomock.Any(), inputParam).Return(nil, errors.New("fake error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			OrderGetterUseCaseMock := ordergetterymock.NewMockOrderGetter(ctrl)

			tt.executeMock(tt.inputParam, OrderGetterUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/orders"

			OrderGetterHandler := ordergetterhandler.NewOrderGetterHandler(OrderGetterUseCaseMock)
			appFake.GET(endpoint, OrderGetterHandler.Handle)

			req, _ := http.NewRequest("GET", fmt.Sprintf("%s?sortBy=%s", endpoint, tt.inputParam), nil)

			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

// temos que verificar aqui o que é o order getter test
