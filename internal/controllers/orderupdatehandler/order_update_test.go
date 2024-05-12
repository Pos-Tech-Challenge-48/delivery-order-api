package orderupdatehandler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/orderupdatehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	orderupdatermock "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/orderupdatormock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderUpdater_Handler(t *testing.T) {

	tests := []struct {
		name               string
		inputBody          *entities.Order
		inputText          string
		expectedStatusCode int
		executeMock        func(m *orderupdatermock.MockOrderUpdater, inputOrder *entities.Order)
	}{
		{
			name: "Should return order from database",
			inputBody: &entities.Order{
				ID:         "fake-order-id",
				CustomerID: "uuid",
				Status:     "Pago",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "uuid",
						Name: "Coca",
					},
					{
						ID:   "uuid2",
						Name: "X-burger",
					},
				}},
			expectedStatusCode: http.StatusOK,
			executeMock: func(m *orderupdatermock.MockOrderUpdater, inputOrder *entities.Order) {
				m.EXPECT().Update(gomock.Any(), inputOrder).Return(nil)
			},
		},
		{
			name: "Update order that doesnt exist",
			inputBody: &entities.Order{
				ID:         "fake-order-id-non-existent",
				CustomerID: "uuid",
				Status:     "Pago",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "uuid",
						Name: "Coca",
					},
					{
						ID:   "uuid2",
						Name: "X-burger",
					},
				}},
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *orderupdatermock.MockOrderUpdater, inputOrder *entities.Order) {
				m.EXPECT().Update(gomock.Any(), inputOrder).Return(entities.ErrOrderNotExist)
			},
		},
		{
			name: "Update caused error",
			inputBody: &entities.Order{
				ID:         "fake-order-id-non-existent",
				CustomerID: "uuid",
				Status:     "Pago",
				OrderProduct: []entities.OrderProduct{
					{
						ID:   "uuid",
						Name: "Coca",
					},
					{
						ID:   "uuid2",
						Name: "X-burger",
					},
				}},
			expectedStatusCode: http.StatusInternalServerError,
			executeMock: func(m *orderupdatermock.MockOrderUpdater, inputOrder *entities.Order) {
				m.EXPECT().Update(gomock.Any(), inputOrder).Return(errors.New("unexpected error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			OrderUpdaterUseCaseMock := orderupdatermock.NewMockOrderUpdater(ctrl)
			tt.executeMock(OrderUpdaterUseCaseMock, tt.inputBody)

			appFake := gin.Default()
			endpoint := "/v1/orders"

			OrderUpdaterHandler := orderupdatehandler.NewOrderUpdaterHandler(OrderUpdaterUseCaseMock)
			appFake.PATCH(fmt.Sprintf("%s/:order_id", endpoint), OrderUpdaterHandler.Handle)

			out, _ := json.Marshal(tt.inputBody)
			req, _ := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", endpoint, tt.inputBody.ID), strings.NewReader(string(out)))

			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}

func Test_OrderUpdaterRequest_Handler(t *testing.T) {

	tests := []struct {
		name               string
		inputText          string
		expectedStatusCode int
		executeMock        func(m *orderupdatermock.MockOrderUpdater)
	}{
		{
			name:               "Body is not an Order",
			inputText:          `{"testenotandorder":"products": [{"id":"uuid","name":"Coca"},{"id":"uuid2","name":"X-burger"}]}`,
			expectedStatusCode: http.StatusBadRequest,
			executeMock: func(m *orderupdatermock.MockOrderUpdater) {
				// endpoint will return before any function call
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			OrderUpdaterUseCaseMock := orderupdatermock.NewMockOrderUpdater(ctrl)
			tt.executeMock(OrderUpdaterUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/orders"

			OrderUpdaterHandler := orderupdatehandler.NewOrderUpdaterHandler(OrderUpdaterUseCaseMock)
			appFake.PATCH(fmt.Sprintf("%s/:order_id", endpoint), OrderUpdaterHandler.Handle)

			req, _ := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", endpoint, "id-fake"), strings.NewReader(tt.inputText))

			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
