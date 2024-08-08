package customerdeletehandler

import (
	"context"
	"errors"
	"net/http"

	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/customerdelete"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customerdelete"
	"github.com/gin-gonic/gin"
)

type CustomerDeleteHandler struct {
	customerDeleteUseCase interfaces.CustomerDelete
}

func NewCustomerDeleteHandler(customerDeleteUseCase interfaces.CustomerDelete) *CustomerDeleteHandler {
	return &CustomerDeleteHandler{
		customerDeleteUseCase: customerDeleteUseCase,
	}
}

func (h *CustomerDeleteHandler) Handle(c *gin.Context) {
	ctx := context.Background()

	ID := c.Param("id")

	err := h.customerDeleteUseCase.Handle(ctx, ID)
	if err != nil {
		if errors.Is(err, customerdelete.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
