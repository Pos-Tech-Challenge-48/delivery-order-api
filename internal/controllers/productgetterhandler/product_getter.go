package productgetterhandler

import (
	"context"
	"net/http"

	_ "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/usecases/productgetter"

	"github.com/gin-gonic/gin"
)

type ProductGetterHandler struct {
	svc productgetter.ProductGetter
}

func NewProductGetterHandler(productGetterUseCase productgetter.ProductGetter) *ProductGetterHandler {
	return &ProductGetterHandler{
		svc: productGetterUseCase,
	}
}

// Products godoc
// @Summary get product by category
// @Param   category     query    string     true        "Category"
// @Tags product
// @Produce application/json
// @Success 200 {array}  entities.Product "Product"
// @Failure 400 {object} string "invalid document"
// @Failure 404 {object} string "customer not find"
// @Failure 500 {object} string "general error"
// @Router /products [get]
func (p *ProductGetterHandler) Handle(c *gin.Context) {
	ctx := context.Background()
	category := c.Request.URL.Query().Get("category")

	list, err := p.svc.GetAll(ctx, category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"response": list})
}
