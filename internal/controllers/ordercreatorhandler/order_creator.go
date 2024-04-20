package ordercreatorhandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/ordercreator"
	"github.com/gin-gonic/gin"
)

type OrderCreatorHandler struct {
	OrderCreatorUseCase interfaces.OrderCreator
}

func NewOrderCreatorHandler(OrderCreatorUseCase interfaces.OrderCreator) *OrderCreatorHandler {
	return &OrderCreatorHandler{
		OrderCreatorUseCase: OrderCreatorUseCase,
	}
}

// CreateOrder godoc
// @Summary create order
// @Description save Order in DB
// @Param Order body entities.Order true "Order"
// @Tags order
// @Produce application/json
// @Success 201
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /orders [post]
func (h *OrderCreatorHandler) Handle(c *gin.Context) {
	var order entities.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	if err := h.OrderCreatorUseCase.CreateOrderAndEnqueuePayment(ctx, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
