package order_enqueuer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/external/sqs_service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type OrderEnqueuerRepository struct {
	queueClient *sqs_service.Client
}

func NewOrderEnqueuerRepository(sqsClient *sqs_service.Client) *OrderEnqueuerRepository {
	return &OrderEnqueuerRepository{
		queueClient: sqsClient,
	}
}

func (c *OrderEnqueuerRepository) FormatOrderMessage(_ context.Context, data *entities.Order) (*OrderMessageDTO, error) {
	return &OrderMessageDTO{
		OrderId:          data.ID,
		OrderStatus:      data.Status,
		Amount:           data.Amount,
		MerchantID:       data.CustomerID,
		CreatedDate:      data.CreatedDate,
		LastModifiedDate: data.LastModifiedDate,
		Email:            "vitorsmap@gmail.com",
	}, nil
}

func (c *OrderEnqueuerRepository) SendPendingPaymentOrderMessageToQueue(ctx context.Context, data *entities.Order) error {
	const operation = "SQS.Enqueue.SendPendingPaymentOrderMessageToQueue"

	queue := c.queueClient.Queues.OrderPaymentConfirmationQueue

	orderMessage, _ := c.FormatOrderMessage(ctx, data)

	body, err := json.Marshal(orderMessage)

	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	_, err = c.queueClient.SQSClient.SendMessage(&sqs.SendMessageInput{
		MessageBody:    aws.String(string(body)),
		QueueUrl:       aws.String(queue.URL),
		MessageGroupId: aws.String("custom-event"),
	})

	if err != nil {
		return fmt.Errorf("%s -> failed to send message: %w", operation, err)
	}

	return nil

}

func (c *OrderEnqueuerRepository) SendOrderToProductionQueue(ctx context.Context, data *entities.Order) error {
	const operation = "SQS.Enqueue.SendOrderToProductionQueue"

	queue := c.queueClient.Queues.OrderProductionQueue

	orderMessage, _ := c.FormatOrderMessage(ctx, data)

	body, err := json.Marshal(orderMessage)

	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	_, err = c.queueClient.SQSClient.SendMessage(&sqs.SendMessageInput{
		MessageBody:    aws.String(string(body)),
		QueueUrl:       aws.String(queue.URL),
		MessageGroupId: aws.String("custom-event"),
	})

	if err != nil {
		return fmt.Errorf("%s -> failed to send message: %w", operation, err)
	}

	return nil
}
