package sqs_service

// Pacote que representa uma conexão com SQS

import (
	"context"
	"fmt"
	"log"
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
	OrderProductionQueue          SQSQueue
}

type Client struct {
	SQSClient *sqs.SQS
	cfg       config.SQSConfig
	Queues    Queues
}

type SQSQueue struct {
	Name    string
	URL     string
	Arn     string
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

// não vai funcionar, o certo aqui vai ser pegar o arn da fila
// TODO: refazer isso daqui

func (c *Client) setupQueues(ctx context.Context, env config.Environment) error {
	queueList := map[string]*SQSQueue{
		"paymentConfimation": {
			Name: c.cfg.QueuePaymentsConfirmation,
		},
		"orderProduction": {
			Name: c.cfg.QueueOrderProduction,
		},
	}

	if env.Name == config.EnvLocal.Name {

		for _, queue := range queueList {
			createQueueInformation, err := c.SQSClient.CreateQueue(&sqs.CreateQueueInput{
				QueueName: aws.String(queue.Name),
				Attributes: map[string]*string{
					"FifoQueue":                 aws.String(strconv.FormatBool(strings.HasSuffix(queue.Name, ".fifo"))),
					"ContentBasedDeduplication": aws.String("true"),
				},
			})

			if err != nil {
				return fmt.Errorf("error creating queue: %w", err)
			}

			queue.URL = *createQueueInformation.QueueUrl
		}

		// manually setup for local environment
		c.Queues.OrderPaymentConfirmationQueue = *queueList["paymentConfimation"]
		c.Queues.OrderProductionQueue = *queueList["orderProduction"]

		return nil
	}

	c.setupOrderPaymentConfirmationQueue()
	c.setupOrderProductionQueue()

	return nil
}

func (c *Client) setupOrderPaymentConfirmationQueue() error {
	orderConfirmationQueue := SQSQueue{
		Name: c.cfg.QueuePaymentsConfirmation,
	}

	log.Printf("setupCustomEventsQueue: %s\n", orderConfirmationQueue.Name)
	customEventsURL, err := c.getQueueURL(orderConfirmationQueue.Name)
	if err != nil {
		log.Printf("%s -> error getting queue url -> %v\n", orderConfirmationQueue.URL, err)
		return fmt.Errorf("error getting queue (%s) url: %w", orderConfirmationQueue.URL, err)
	}

	orderConfirmationQueue.URL = *customEventsURL
	orderConfirmationQueue.Arn = getQueueArn(*customEventsURL)
	c.Queues.OrderPaymentConfirmationQueue = orderConfirmationQueue

	return nil
}

func (c *Client) setupOrderProductionQueue() error {
	orderProductionQueue := SQSQueue{
		Name: c.cfg.QueueOrderProduction,
	}

	log.Printf("setup order production queue: %s\n", orderProductionQueue.Name)
	customEventsURL, err := c.getQueueURL(orderProductionQueue.Name)
	if err != nil {
		log.Printf("%s -> error getting queue url -> %v\n", orderProductionQueue.URL, err)
		return fmt.Errorf("error getting queue (%s) url: %w", orderProductionQueue.URL, err)
	}

	orderProductionQueue.URL = *customEventsURL
	orderProductionQueue.Arn = getQueueArn(*customEventsURL)
	c.Queues.OrderProductionQueue = orderProductionQueue

	return nil
}

func (c *Client) getQueueURL(name string) (*string, error) {

	output, err := c.SQSClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		log.Printf("error getting queue URL -> %v\n", err)
		return nil, fmt.Errorf("error getting queue URL -> %w", name, err)
	}

	return output.QueueUrl, nil
}

func getQueueArn(queueURL string) string {
	log.Printf("To get queue arn: %s\n", queueURL)
	parts := strings.Split(queueURL, "/")
	subParts := strings.Split(parts[2], ".")
	arn := "arn:aws:" + subParts[0] + ":" + subParts[1] + ":" + parts[3] + ":" + parts[4]
	log.Printf("Got queue arn: %s\n", arn)
	return arn
}
