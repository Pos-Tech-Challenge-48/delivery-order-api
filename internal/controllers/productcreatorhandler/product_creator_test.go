package productcreatorhandler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/productcreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/mocks/productcreatormock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_ProductCreator_Handler(t *testing.T) {

	fakeProduct := &entities.Product{
		Name:        "Coca-Cola Zero",
		CategoryID:  "Bebida",
		Description: "Uma coquinha geladinha.",
		Price:       13.50,
		Image:       "https://unsplash.com/photos/filled-coca-cola-bottle-XWdRIu-Rk_0",
	}

	tests := []struct {
		name               string
		inputBody          string
		expectedStatusCode int
		executeMock        func(m *productcreatormock.MockProductCreator)
	}{
		{
			name: "Should Product created with success",
			inputBody: `{
				"category": "Bebida",
				"description": "Uma coquinha geladinha.",
				"image": "https://unsplash.com/photos/filled-coca-cola-bottle-XWdRIu-Rk_0",
				"name": "Coca-Cola Zero",
				"price": 13.50
			  }`,
			expectedStatusCode: http.StatusOK,
			executeMock: func(m *productcreatormock.MockProductCreator) {
				m.EXPECT().Add(gomock.Any(), fakeProduct).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			ProductCreatorUseCaseMock := productcreatormock.NewMockProductCreator(ctrl)

			tt.executeMock(ProductCreatorUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/products"

			ProductCreatorHandler := productcreatorhandler.NewProductCreatorHandler(ProductCreatorUseCaseMock)

			appFake.POST(endpoint, ProductCreatorHandler.Handle)
			req, _ := http.NewRequest("POST", endpoint, strings.NewReader(tt.inputBody))
			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
