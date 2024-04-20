package customergetterhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/customergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customergetter"
	"github.com/gin-gonic/gin"
)

type CustomerGetterHandler struct {
	customerGetterUseCase interfaces.CustomerGetter
}

func NewCustomerGetterHandler(customerGetterUseCase interfaces.CustomerGetter) *CustomerGetterHandler {
	return &CustomerGetterHandler{
		customerGetterUseCase: customerGetterUseCase,
	}
}

// CreateCustomer godoc
// @Summary get customer by document
// @Param   document     query    string     true        "Document"
// @Tags customer
// @Produce application/json
// @Success 200 {object} entities.Customer "Customer"
// @Failure 400 {object} string "invalid document"
// @Failure 404 {object} string "customer not find"
// @Failure 500 {object} string "general error"
// @Router /customer [get]
func (h *CustomerGetterHandler) Handle(c *gin.Context) {

	ctx := context.Background()

	document := c.Query("document")

	customerInput := entities.Customer{
		Document: document,
	}

	customer, err := h.customerGetterUseCase.Get(ctx, &customerInput)
	if err != nil {
		if errors.Is(err, customergetter.ErrCustomerNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, entities.ErrCustomerInvalidDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, entities.ErrCustomerEmptyDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)

}
