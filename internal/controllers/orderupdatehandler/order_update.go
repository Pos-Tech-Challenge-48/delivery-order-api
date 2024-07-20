package orderupdatehandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/orderupdater"

	"github.com/gin-gonic/gin"
)

type OrderUpdaterHandler struct {
	OrderUpdaterUseCase interfaces.OrderUpdater
}

func NewOrderUpdaterHandler(OrderUpdaterUseCase interfaces.OrderUpdater) *OrderUpdaterHandler {
	return &OrderUpdaterHandler{
		OrderUpdaterUseCase: OrderUpdaterUseCase,
	}
}

// Order Update godoc
// @Summary update Order
// @Description update Order in DB
// @Param order_id path string true "Order ID"
// @Param Order body entities.Order true "Order"
// @Tags order
// @Produce application/json
// @Success 200
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /orders/{order_id} [patch]
func (h *OrderUpdaterHandler) Handle(c *gin.Context) {
	orderID := c.Param("order_id")

	ctx := context.Background()

	var orderBody *entities.Order

	if err := c.BindJSON(&orderBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	orderBody.ID = orderID

	err := h.OrderUpdaterUseCase.Update(ctx, orderBody)
	if err != nil {
		if errors.Is(err, entities.ErrOrderNotExist) || errors.Is(err, entities.ErrInvalidOrderStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
