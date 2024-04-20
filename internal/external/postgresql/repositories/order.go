package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	_ "github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *entities.Order) error {
	query := `
	INSERT INTO restaurant_order (restaurant_order_id, restaurant_order_customer_id, restaurant_order_status_id, restaurant_order_amount)
	VALUES ($1, $2, (SELECT status_id from status where status_name = $3), $4)
`

	_, err := r.db.Exec(
		query,
		order.ID,
		order.CustomerID,
		order.Status,
		order.Amount)

	if err != nil {
		return err
	}

	for _, p := range order.OrderProduct {

		query := `
        INSERT INTO order_item (order_item_product_id, order_item_order_id)
        VALUES ($1, $2)
    `
		_, err := r.db.Exec(query, p.ID, order.ID)
		if err != nil {
			return err
		}

	}

	return err
}

func (r *OrderRepository) GetAll(ctx context.Context) ([]entities.Order, error) {
	queryParams := []interface{}{}

	query := `
	SELECT
		restaurant_order_id,
		restaurant_order_customer_id,
		status.status_name,
		restaurant_order_amount,
		restaurant_order.created_date_db,
		restaurant_order.last_modified_date_db,
		json_agg(json_build_object(
				'id',     order_item.order_item_id,
				'name', p.product_name
			)) AS order_items
	FROM restaurant_order
			join order_item on restaurant_order.restaurant_order_id = order_item.order_item_order_id
			join product p on order_item.order_item_product_id = p.product_id
			JOIN status ON	restaurant_order_status_id = status_id
	GROUP BY
		restaurant_order_id,
		restaurant_order_customer_id,
		status.status_name,
		restaurant_order_amount,
		restaurant_order.created_date_db,
		restaurant_order.last_modified_date_db;
	`

	rows, err := r.db.Query(query, queryParams...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]entities.Order, 0)
	var order entities.Order

	var orderProducts []byte

	for rows.Next() {
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Status,
			&order.Amount,
			&order.CreatedDate,
			&order.LastModifiedDate,
			&orderProducts,
		)

		if err != nil {
			return result, err
		}

		err = json.Unmarshal(orderProducts, &order.OrderProduct)
		if err != nil {
			return result, err
		}

		result = append(result, order)

	}

	return result, nil

}

func (r *OrderRepository) GetAllSortedByStatus(ctx context.Context) ([]entities.Order, error) {
	query := `
    SELECT
        restaurant_order_id,
        status.status_name
    FROM restaurant_order
        JOIN status ON restaurant_order_status_id = status_id
    WHERE status.status_name != 'Finalizado'
    ORDER BY 
        CASE status.status_name
            WHEN 'Pronto' THEN 1
            WHEN 'Em Preparação' THEN 2
            WHEN 'Recebido' THEN 3
        END,
        restaurant_order.created_date_db ASC;
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		err := rows.Scan(
			&order.ID,
			&order.Status,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, orderID string) (*entities.Order, error) {
	query := `
		SELECT ro.restaurant_order_id, ro.restaurant_order_customer_id, s.status_name, ro.restaurant_order_amount, ro.created_date_db, ro.last_modified_date_db 
		FROM restaurant_order ro
		INNER JOIN status s ON ro.restaurant_order_status_id = s.status_id WHERE ro.restaurant_order_id = $1
	`

	row := r.db.QueryRowContext(ctx, query, orderID)

	order := &entities.Order{}
	err := row.Scan(
		&order.ID,
		&order.CustomerID,
		&order.Status,
		&order.Amount,
		&order.CreatedDate,
		&order.LastModifiedDate,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *entities.Order) error {
	query := `
		UPDATE restaurant_order
		SET restaurant_order_customer_id = $1,
			restaurant_order_status_id = (SELECT status_id FROM status WHERE status_name = $2),
			restaurant_order_amount = $3,
			last_modified_date_db = NOW()
		WHERE restaurant_order_id = $4
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		order.CustomerID,
		order.Status,
		order.Amount,
		order.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, order *entities.Order) error {
	query := `
		UPDATE restaurant_order
		SET restaurant_order_status_id = (SELECT status_id FROM status WHERE status_name = $1),
			last_modified_date_db = NOW()
		WHERE restaurant_order_id = $2
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		order.Status,
		order.ID,
	)

	if err != nil {
		return err
	}
	return nil
}
