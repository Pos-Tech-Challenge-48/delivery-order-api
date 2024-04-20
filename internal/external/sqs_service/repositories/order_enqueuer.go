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

func (c *OrderEnqueuerRepository) SendPendingPaymentOrderMessageToQueue(ctx context.Context, data *entities.Order) error {
	const operation = "SQS.Enqueue.SendMessageToQueue"

	queue := c.queueClient.Queues.OrderPaymentConfirmationQueue

	body, err := json.Marshal(data)

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
