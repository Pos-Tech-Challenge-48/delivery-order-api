package controllers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
	CustomerGetterHandler  gin.HandlerFunc
	CustomerDeleteHandler  gin.HandlerFunc
	OrderCreatorHandler    gin.HandlerFunc
	OrderGetterHandler     gin.HandlerFunc
	OrderUpdaterHandler    gin.HandlerFunc
	ProductCreatorHandler  gin.HandlerFunc
	ProductDeleteHandler   gin.HandlerFunc
	ProductUpdateHandler   gin.HandlerFunc
	ProductGetterHandler   gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	delivery := app.Group("/v1/delivery")
	{

		// CUSTOMERS
		delivery.POST("/customers", r.CustomerCreatorHandler)
		delivery.GET("/customers", r.CustomerGetterHandler)
		delivery.DELETE("/customers/:id", r.CustomerDeleteHandler)

		// product routes
		delivery.POST("/products", r.ProductCreatorHandler)
		delivery.PUT("/products", r.ProductUpdateHandler)
		delivery.DELETE("/products", r.ProductDeleteHandler)
		delivery.GET("/products", r.ProductGetterHandler)

		// order routes
		delivery.POST("/orders", r.OrderCreatorHandler)
		delivery.GET("/orders", r.OrderGetterHandler)
		delivery.PATCH("/orders/:order_id", r.OrderUpdaterHandler)
	}
}
