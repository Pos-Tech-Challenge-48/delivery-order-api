package repositories

import (
	"context"
	"database/sql"

	"github.com/Pos-Tech-Challenge-48/delivery-order-api/internal/entities"
	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Save(ctx context.Context, customer *entities.Customer) error {
	query := `
	INSERT INTO customer (customer_id, customer_name, customer_email, customer_document, created_date_db, last_modified_date_db)
	VALUES ($1, $2, $3, $4, $5, $6)
`
	_, err := r.db.Exec(
		query,
		customer.ID,
		customer.Name,
		customer.Email,
		customer.Document,
		customer.CreatedDate,
		customer.LastModifiedDate,
	)
	return err
}

func (r *CustomerRepository) GetByDocument(ctx context.Context, document string) (*entities.Customer, error) {
	query := `
        SELECT customer_id, customer_name, customer_email, customer_document, created_date_db, last_modified_date_db
        FROM customer
        WHERE customer_document = $1
    `
	row := r.db.QueryRow(query, document)

	customer := &entities.Customer{}
	err := row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Document,
		&customer.CreatedDate,
		&customer.LastModifiedDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) GetByDocumentAndEmail(ctx context.Context, document string, email string) (*entities.Customer, error) {
	query := `
        SELECT customer_id, customer_name, customer_email, customer_document, created_date_db, last_modified_date_db
        FROM customer
        WHERE customer_document = $1 AND customer_email = $2
    `
	row := r.db.QueryRow(query, document, email)

	customer := &entities.Customer{}
	err := row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Document,
		&customer.CreatedDate,
		&customer.LastModifiedDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, ID string) (*entities.Customer, error) {
	query := `
        SELECT customer_id, customer_name, customer_email, customer_document, created_date_db, last_modified_date_db
        FROM customer
        WHERE customer_id = $1
    `
	row := r.db.QueryRow(query, ID)

	customer := &entities.Customer{}
	err := row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Document,
		&customer.CreatedDate,
		&customer.LastModifiedDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) Delete(ctx context.Context, ID string) error {
	query := `
        DELETE FROM customer
        WHERE customer_id = $1
    `
	_, err := r.db.Exec(query, ID)
	return err
}
