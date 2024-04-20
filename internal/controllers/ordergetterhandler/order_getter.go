package ordergetterhandler

import (
	"context"
	"net/http"

	_ "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/ordergetter"
	"github.com/gin-gonic/gin"
)

type OrderGetterHandler struct {
	OrderGetterUseCase interfaces.OrderGetter
}

func NewOrderGetterHandler(OrderGetterUseCase interfaces.OrderGetter) *OrderGetterHandler {
	return &OrderGetterHandler{
		OrderGetterUseCase: OrderGetterUseCase,
	}
}

// GetOrder godoc
// @Summary get all order
// @Description Get Order from DB
// @Tags order
// @Produce application/json
// @Param sortBy query string false "Sort orders by status"
// @Success 200 {array} entities.Order "Order"
// @Failure 400 {object} string "invalid document"
// @Failure 404 {object} string "customer not find"
// @Failure 500 {object} string "general error"
// @Router /orders [get]
func (o *OrderGetterHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	sortBy := c.Query("sortBy")

	list, err := o.OrderGetterUseCase.GetAll(ctx, sortBy)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": list})
}
