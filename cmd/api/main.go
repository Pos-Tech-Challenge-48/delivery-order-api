package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/customercreatorhandler"
	customergetterhandler "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/customergetterandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/ordercreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/ordergetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/orderupdatehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/productcreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/productdeletehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/productgetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/controllers/productupdatehandler"
	db "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/postgresql"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/postgresql/repositories"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/sqs_service"
	enqueuer_repository "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/sqs_service/repositories"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/customergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/ordercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/ordergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/orderupdater"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productcreator"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productdelete"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productgetter"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/usecases/productupdate"

	"github.com/gin-gonic/gin"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/config"
)

func main() {
	mainCtx := context.Background()

	config, err := config.LoadConfig()
	if err != nil {
		log.Println("error to load config", err)

	}
	// SQS
	sqsService, err := sqs_service.New(mainCtx, config.SQSConfig, config.Environment)
	if err != nil {
		log.Println("error: failed to start SQS %w", err)
	}

	// POSTGRESQL
	postgres := db.New(config)
	defer postgres.Close()

	// CUSTOMER REPOSITORY
	customerRepository := repositories.NewCustomerRepository(postgres)
	customerCreator := customercreator.NewCustomerCreator(customerRepository)
	customerCreatorHandler := customercreatorhandler.NewCustomerCreatorHandler(customerCreator)

	customerGetter := customergetter.NewCustomerGetter(customerRepository)
	customerGetterHandler := customergetterhandler.NewCustomerGetterHandler(customerGetter)

	// PRODUCTS
	productRepository := repositories.NewProductRepository(postgres)

	productCreator := productcreator.NewProductCreator(productRepository)
	productCreatorHandler := productcreatorhandler.NewProductCreatorHandler(productCreator)

	productDelete := productdelete.NewProductDelete(productRepository)
	productDeleteHandler := productdeletehandler.NewProductDeleteHandler(productDelete)

	productGetter := productgetter.NewProductGetter(productRepository)
	productGetterHandler := productgetterhandler.NewProductGetterHandler(productGetter)

	productUpdate := productupdate.NewProductUpdate(productRepository)
	productUpdateHandler := productupdatehandler.NewProductUpdateHandler(productUpdate)

	// ORDER REPOSITORY
	orderRepository := repositories.NewOrderRepository(postgres)

	orderEnqueuerRepository := enqueuer_repository.NewOrderEnqueuerRepository(sqsService)

	orderCreator := ordercreator.NewOrderCreator(orderRepository, productRepository, orderEnqueuerRepository)
	orderCreatorHandler := ordercreatorhandler.NewOrderCreatorHandler(orderCreator)

	orderGetter := ordergetter.NewOrderGetter(orderRepository, productRepository)
	orderGetterHandler := ordergetterhandler.NewOrderGetterHandler(orderGetter)

	orderUpdater := orderupdater.NewOrderUpdater(orderRepository, orderEnqueuerRepository)
	orderUpdaterHandler := orderupdatehandler.NewOrderUpdaterHandler(orderUpdater)

	app := gin.Default()

	router := controllers.Router{
		CustomerCreatorHandler: customerCreatorHandler.Handle,
		CustomerGetterHandler:  customerGetterHandler.Handle,
		OrderCreatorHandler:    orderCreatorHandler.Handle,
		OrderGetterHandler:     orderGetterHandler.Handle,
		OrderUpdaterHandler:    orderUpdaterHandler.Handle,
		ProductCreatorHandler:  productCreatorHandler.Handle,
		ProductDeleteHandler:   productDeleteHandler.Handle,
		ProductUpdateHandler:   productUpdateHandler.Handle,
		ProductGetterHandler:   productGetterHandler.Handle,
	}

	router.Register(app)

	app.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	app.Run(":8081")
}
