package customercreatorhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/customercreator"
	"github.com/gin-gonic/gin"
)

type CustomerCreatorHandler struct {
	customerCreatorUseCase customercreator.CustomerCreator
}

func NewCustomerCreatorHandler(customerCreatorUseCase customercreator.CustomerCreator) *CustomerCreatorHandler {
	return &CustomerCreatorHandler{
		customerCreatorUseCase: customerCreatorUseCase,
	}
}

// CreateCustomer godoc
// @Summary create customer
// @Description save customer in DB
// @Param customer body entities.Customer true "Customer"
// @Tags customer
// @Produce application/json
// @Success 201
// @Failure 400 {string} message  "invalid document, invalid email..."
// @Failure 500 {string} message  "general error"
// @Router /customer [post]
func (h *CustomerCreatorHandler) Handle(c *gin.Context) {
	var customer entities.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	if err := h.customerCreatorUseCase.Create(ctx, &customer); err != nil {
		if errors.Is(err, entities.ErrCustomerEmptyEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, entities.ErrCustomerEmptyName) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, customercreator.ErrAlreadyExistsCustomer) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, entities.ErrCustomerInvalidDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, entities.ErrCustomerInvalidEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
