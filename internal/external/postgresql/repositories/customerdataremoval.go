package repositories

import (
	"context"
	"database/sql"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

type CustomerDataRemovalRepository struct {
	db *sql.DB
}

func NewCustomerDataRemovalRepository(db *sql.DB) *CustomerDataRemovalRepository {
	return &CustomerDataRemovalRepository{
		db: db,
	}
}

func (r *CustomerDataRemovalRepository) Save(ctx context.Context, data *entities.CustomerDataRemovalRequest) error {
	query := `
	INSERT INTO customer_data_removal_requests (request_id, requester_name, requester_address, requester_phonenumber, created_date_db, last_modified_date_db)
	VALUES ($1, $2, $3, $4, $5, $6)
`
	_, err := r.db.Exec(
		query,
		data.ID,
		data.Name,
		data.Address,
		data.Phonenumber,
		data.CreatedDate,
		data.LastModifiedDate,
	)
	return err
}
