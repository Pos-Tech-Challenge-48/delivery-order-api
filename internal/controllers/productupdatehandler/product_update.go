package productupdatehandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/productupdate"
	"github.com/gin-gonic/gin"
)

type ProductUpdateHandler struct {
	svc interfaces.ProductUpdate
}

func NewProductUpdateHandler(productUpdateUseCase interfaces.ProductUpdate) *ProductUpdateHandler {
	return &ProductUpdateHandler{
		svc: productUpdateUseCase,
	}
}

// Products godoc
// @Summary update product
// @Description update product in DB
// @Param product body entities.Product true "Product"
// @Tags product
// @Produce application/json
// @Success 200
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /products [put]
func (p *ProductUpdateHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	var product entities.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	err = product.Validate()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	err = p.svc.Update(ctx, &product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Product updated successfully"})
}
