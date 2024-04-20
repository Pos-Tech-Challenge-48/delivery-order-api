package productcreatorhandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/productcreator"
	"github.com/gin-gonic/gin"
)

type ProductCreatorHandler struct {
	svc interfaces.ProductCreator
}

func NewProductCreatorHandler(productAddUseCase interfaces.ProductCreator) *ProductCreatorHandler {
	return &ProductCreatorHandler{
		svc: productAddUseCase,
	}
}

// Products godoc
// @Summary create product
// @Description save product in DB
// @Param product body entities.Product true "Product"
// @Tags product
// @Produce application/json
// @Success 200
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /products [post]
func (p *ProductCreatorHandler) Handle(c *gin.Context) {
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

	err = p.svc.Add(ctx, &product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Product created successfully"})
}
