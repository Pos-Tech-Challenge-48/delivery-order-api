package sqs_service

// Pacote que representa uma conexão com SQS

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Queues struct {
	OrderPaymentConfirmationQueue SQSQueue
}

type Client struct {
	SQSClient *sqs.SQS
	cfg       config.SQSConfig
	Queues    Queues
}

type SQSQueue struct {
	Name    string
	URL     string
	Handler Handler
}

func New(ctx context.Context, cfg config.SQSConfig, env config.Environment) (*Client, error) {
	const operation = "SQS.New"

	awsConfig := &aws.Config{MaxRetries: aws.Int(5)}

	if env.Name == config.EnvLocal.Name {
		awsConfig.Region = aws.String("us-east-1")
		awsConfig.Endpoint = aws.String(cfg.LocalURL)
		awsConfig.Credentials = credentials.NewStaticCredentials("test", "test", "")
	}

	sess, err := session.NewSession(awsConfig)

	if err != nil {
		return nil, err
	}

	sqsClient := sqs.New(sess)

	client := &Client{sqsClient, cfg, Queues{}}

	err = client.setupQueues()
	if err != nil {
		return nil, fmt.Errorf("error %s: %w", operation, err)
	}

	return client, nil
}

func (c *Client) setupQueues() error {

	// Payments Confirmation Queue
	paymentConfirmationQueue := &SQSQueue{
		Name: c.cfg.QueuePaymentsConfirmation,
	}

	createQueueInformation, err := c.SQSClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(paymentConfirmationQueue.Name),
		Attributes: map[string]*string{
			"FifoQueue":                 aws.String(strconv.FormatBool(strings.HasSuffix(paymentConfirmationQueue.Name, ".fifo"))),
			"ContentBasedDeduplication": aws.String("true"),
		},
	})

	if err != nil {
		return fmt.Errorf("error creating queue: %w", err)
	}

	paymentConfirmationQueue.URL = *createQueueInformation.QueueUrl
	c.Queues.OrderPaymentConfirmationQueue = *paymentConfirmationQueue

	return nil
}
