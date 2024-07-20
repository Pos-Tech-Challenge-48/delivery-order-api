package customerdataremovalhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	customerdataremovalcreator "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/customerdararemovalcreator"
	"github.com/gin-gonic/gin"
)

type CustomerDataRemovalRequestHandler struct {
	customerDataRemovalUseCase customerdataremovalcreator.CustomerDataRemovalRequestCreator
}

func NewCustomerDataRemovalRequestHandler(customerDataRemovalUseCase customerdataremovalcreator.CustomerDataRemovalRequestCreator) *CustomerDataRemovalRequestHandler {
	return &CustomerDataRemovalRequestHandler{
		customerDataRemovalUseCase: customerDataRemovalUseCase,
	}
}

// Create Customer Data Removal Request godoc
// @Summary create customerdataremoval
// @Description save customerdataremoval in DB
// @Param customerdataremoval body entities.CustomerDataRemovalRequest true "CustomerDataRemovalRequest"
// @Tags customer
// @Produce application/json
// @Success 201
// @Failure 400 {string} message  "invalid customer data removal"
// @Failure 500 {string} message  "general error"
// @Router /customers/data_removal_request [post]
func (h *CustomerDataRemovalRequestHandler) Handle(c *gin.Context) {
	var data entities.CustomerDataRemovalRequest

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	if err := h.customerDataRemovalUseCase.Create(ctx, &data); err != nil {
		if errors.Is(err, entities.ErrCustomerDataRemoval) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
