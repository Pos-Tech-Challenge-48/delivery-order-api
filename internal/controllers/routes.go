package controllers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	delivery := app.Group("/v1/delivery")
	{
		delivery.POST("/customers", r.CustomerCreatorHandler)
	}
}
