package orderupdater

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/interfaces/repositories"
)

type OrderUpdater struct {
	orderRepository interfaces.OrderRepository
	orderEnqueuer   interfaces.OrderQueueRepository
}

func NewOrderUpdater(orderRepository interfaces.OrderRepository, orderEnqueuer interfaces.OrderQueueRepository) *OrderUpdater {
	return &OrderUpdater{
		orderRepository: orderRepository,
		orderEnqueuer:   orderEnqueuer,
	}
}

func (p *OrderUpdater) Update(ctx context.Context, order *entities.Order) error {
	existingOrder, err := p.orderRepository.GetByID(ctx, order.ID)
	if err != nil {
		return err
	}

	err = p.validateOrderForUpdating(existingOrder)
	if err != nil {
		return err
	}

	existingOrder.Status = order.Status

	err = p.orderRepository.Update(ctx, existingOrder)
	if err != nil {
		return err
	}

	shouldEnqueue := p.validateForEnqueueMessage(existingOrder)
	if shouldEnqueue {
		p.orderEnqueuer.SendOrderToProductionQueue(ctx, existingOrder)
	}

	return nil

}

func (p *OrderUpdater) validateForEnqueueMessage(order *entities.Order) bool {
	return order.Status == "Pago"
}

func (p *OrderUpdater) validateOrderForUpdating(order *entities.Order) error {

	if order == nil {
		return entities.ErrOrderNotExist
	}

	if order.IsFinished() {
		return entities.ErrInvalidOrderStatus
	}

	return nil
}
