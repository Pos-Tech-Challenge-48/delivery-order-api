package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrOrderNotExist                 = errors.New("order not exists")
	ErrOrderNotReady                 = errors.New("order not ready")
	ErrInvalidOrderStatus            = errors.New("status not able to be updated")
	ErrOrderNotReadyToConfirmPayment = errors.New("order not ready to confirm the payment")
)

type Order struct {
	ID               string         `json:"id"`
	CustomerID       string         `json:"customer_id"`
	Status           string         `json:"status"`
	Amount           float64        `json:"amount"`
	OrderProduct     []OrderProduct `json:"products"`
	CreatedDate      time.Time      `json:"created_date_db"`
	LastModifiedDate time.Time      `json:"last_modified_date_db"`
}

type OrderProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewOrder(customerID string, products []OrderProduct) *Order {
	return &Order{
		ID:               uuid.NewString(),
		CustomerID:       customerID,
		Status:           "Recebido",
		OrderProduct:     products,
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
	}
}

func (o *Order) Validate() error {
	return nil
}

func (o *Order) IsReadyToPayment() bool {
	return o.Status == "Recebido"
}

func (o *Order) IsFinished() bool {
	return o.Status == "Finalizado"
}

func (o *Order) IsReadyToConfirmPayment() bool {
	return o.Status == "Pendente Pagamento"
}

func (o *Order) SetPendingPayment() *Order {
	o.Status = "Pendente Pagamento"
	return o
}

func (o *Order) SetPaid() *Order {
	o.Status = "Pago"
	return o
}

func (o *Order) SetInProgress() *Order {
	o.Status = "Em Preparação"
	return o
}

func (o *Order) SetReady() *Order {
	o.Status = "Pronto"
	return o
}

func (o *Order) SetFinished() *Order {
	o.Status = "Finalizado"
	return o
}

func (o *Order) SetPaymentRefused() *Order {
	o.Status = "Pagamento Recusado"
	return o
}
