package ordercreator

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

type OrderCreator struct {
	orderRepository   interfaces.OrderRepository
	productRepository interfaces.ProductRepository
	orderEnqueuer     interfaces.OrderQueueRepository
}

func NewOrderCreator(orderRepository interfaces.OrderRepository, productRepository interfaces.ProductRepository, orderEnqueuer interfaces.OrderQueueRepository) *OrderCreator {
	return &OrderCreator{
		orderRepository:   orderRepository,
		productRepository: productRepository,
		orderEnqueuer:     orderEnqueuer,
	}
}

func (uc *OrderCreator) CreateOrderAndEnqueuePayment(ctx context.Context, orderInput *entities.Order) error {
	err := orderInput.Validate()
	if err != nil {
		return err
	}

	order := entities.NewOrder(orderInput.CustomerID, orderInput.OrderProduct)

	amount, err := uc.calculateAmount(ctx, order.OrderProduct)
	if err != nil {
		return err
	}

	order.Amount = amount

	err = uc.orderRepository.Save(ctx, order)
	if err != nil {
		return err
	}

	return uc.orderEnqueuer.SendPendingPaymentOrderMessageToQueue(ctx, order)
}

func (uc *OrderCreator) calculateAmount(ctx context.Context, products []entities.OrderProduct) (float64, error) {
	var amount float64

	for _, p := range products {

		product, err := uc.productRepository.GetByID(ctx, p.ID)
		if err != nil {
			return 0, err
		}
		amount += product.Price

	}
	return amount, nil

}

// vai ter que existir o product tbm
