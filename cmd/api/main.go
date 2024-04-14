package main

import (
	"log"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/customercreatorhandler"
	db "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/postgresql"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/postgresql/repositories"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customercreator"

	"github.com/gin-gonic/gin"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Panic("error to load config", err)
	}

	// POSTGRESQL
	postgres := db.New(config)
	defer postgres.Close()

	customerRepository := repositories.NewCustomerRepository(postgres)

	customerCreator := customercreator.NewCustomerCreator(customerRepository)
	customerCreatorHandler := customercreatorhandler.NewCustomerCreatorHandler(customerCreator)

	app := gin.Default()

	router := controllers.Router{
		CustomerCreatorHandler: customerCreatorHandler.Handle,
	}

	router.Register(app)

	app.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	app.Run(":8081")
}
