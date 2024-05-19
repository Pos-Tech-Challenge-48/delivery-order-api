package order_enqueuer

import "time"

type OrderMessageDTO struct {
	OrderId     string  `json:"order_id"`
	OrderStatus string  `json:"status"`
	Amount      float64 `json:"amount"`
	MerchantID  string  `json:"merchant_id"`
	Email       string  `json:"email"`

	CreatedDate      time.Time `json:"created_date_db"`
	LastModifiedDate time.Time `json:"last_modified_date_db"`
}
