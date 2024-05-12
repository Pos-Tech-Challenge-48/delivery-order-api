package productdeletehandler

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productdelete"
	"github.com/gin-gonic/gin"
)

type ProductDeleteHandler struct {
	svc productdelete.ProductDelete
}

func NewProductDeleteHandler(productDeleteUseCase *productdelete.ProductDelete) *ProductDeleteHandler {
	return &ProductDeleteHandler{
		svc: *productDeleteUseCase,
	}
}

// Products godoc
// @Summary delete product
// @Description delete product in DB
// @Param product body entities.Product true "Product"
// @Tags product
// @Produce application/json
// @Success 200
// @Failure 400 {string} message  "invalid request"
// @Failure 500 {string} message  "general error"
// @Router /products [delete]
func (p *ProductDeleteHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	var product entities.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	err = p.svc.Delete(ctx, product.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "Product deleted successfully"})
}
