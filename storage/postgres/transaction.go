package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"teamProject/api/models"
	"teamProject/storage"
)

type transactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) storage.ITransactionStorage {
	return transactionRepo{db: db}
}

func (t transactionRepo) Create(trans models.CreateTransaction) (string, error) {
	id := uuid.New()
	query := `insert into transactions 
    					(id, sale_id, staff_id, transaction_type, source_type, from_amount, to_amount, description) 
						values ($1, $2, $3, $4, $5, $6, $7, $8)`
	if _, err := t.db.Exec(context.Background(), query, id,
		trans.SaleID,
		trans.StaffID,
		trans.TransactionType,
		trans.SourceType,
		trans.FromAmount,
		trans.ToAmount,
		trans.Description); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return "", err
	}
	return id.String(), nil
}

func (t transactionRepo) GetByID(id string) (models.Transaction, error) {
	trans := models.Transaction{}
	query := `select id, sale_id, staff_id, transaction_type, source_type, from_amount, to_amount,
       						description, created_at, updated_at
							from transactions where deleted_at is null and id = $1`
	if err := t.db.QueryRow(context.Background(), query, id).Scan(
		&trans.ID,
		&trans.SaleID,
		&trans.StaffID,
		&trans.TransactionType,
		&trans.SourceType,
		&trans.FromAmount,
		&trans.ToAmount,
		&trans.Description,
		&trans.CreatedAt,
		&trans.UpdatedAt); err != nil {
		fmt.Println("error is while selecting by id", err.Error())
		return models.Transaction{}, err
	}
	return trans, nil
}

func (t transactionRepo) GetList(request models.GetListRequest) (models.TransactionResponse, error) {
	var (
		page              = request.Page
		offset            = (page - 1) * request.Limit
		transactions      = []models.Transaction{}
		search            = request.Search
		count             = 0
		query, countQuery string
	)

	countQuery = `select count(1) from transactions where deleted_at is null `
	if search != "" {
		countQuery += fmt.Sprintf(` and CAST(from_amount AS TEXT) ilike '%%%s%%' or 
											CAST(to_amount AS TEXT) ilike '%%%s%%'`, search, search)
	}
	if err := t.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning row", err.Error())
		return models.TransactionResponse{}, err
	}

	query = `select id, sale_id, staff_id, transaction_type, source_type, from_amount, to_amount,
       						description, created_at, updated_at from transactions where deleted_at is null `
	if search != "" {
		query += fmt.Sprintf(` and CAST(from_amount AS TEXT) ilike '%%%s%%' or 
											CAST(to_amount AS TEXT) ilike '%%%s%%' `, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := t.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting all from transactions", err.Error())
		return models.TransactionResponse{}, err
	}

	for rows.Next() {
		trans := models.Transaction{}
		if err = rows.Scan(
			&trans.ID,
			&trans.SaleID,
			&trans.StaffID,
			&trans.TransactionType,
			&trans.SourceType,
			&trans.FromAmount,
			&trans.ToAmount,
			&trans.Description,
			&trans.CreatedAt,
			&trans.UpdatedAt); err != nil {
			fmt.Println("error is while scanning rows", err.Error())
			return models.TransactionResponse{}, err
		}
		transactions = append(transactions, trans)
	}
	return models.TransactionResponse{
		Transactions: transactions,
		Count:        count,
	}, nil
}

func (t transactionRepo) Update(transaction models.UpdateTransaction) (string, error) {
	query := `update transactions set sale_id = $1, staff_id = $2, transaction_type = $3, source_type = $4,from_amount = $5, to_amount = $6,
								description = $7, updated_at = now() 
                    			where id = $8`
	if _, err := t.db.Exec(context.Background(), query,
		&transaction.SaleID,
		&transaction.StaffID,
		&transaction.TransactionType,
		&transaction.SourceType,
		&transaction.FromAmount,
		&transaction.ToAmount,
		&transaction.Description,
		&transaction.ID); err != nil {
		fmt.Println("error is while updating transaction", err.Error())
		return "", err
	}
	return transaction.ID, nil
}

func (t transactionRepo) Delete(id string) error {
	query := `update transactions set deleted_at = now() where id = $1`
	if _, err := t.db.Exec(context.Background(), query, id); err != nil {
		fmt.Println("error is while deleting", err.Error())
		return err
	}
	return nil
}
