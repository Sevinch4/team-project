package postgres

import (
	"context"
	"fmt"
	"log"
	"teamProject/api/models"
	"teamProject/storage"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repositoryTransactionRepo struct {
	DB *pgxpool.Pool
}

func NewRepositoryTransactionRepo(DB *pgxpool.Pool) storage.IRepositoryTransactionRepo {
	return &repositoryTransactionRepo{
		DB: DB,
	}
}

func (s *repositoryTransactionRepo) Create(rtransaction models.CreateRepositoryTransaction) (string, error) {
	id := uuid.New().String()
	createdAt := time.Now()

	if _, err := s.DB.Exec(context.Background(), `INSERT INTO repository_transactions
		(id, staff_id, product_id, repository_transaction_type, price, quantity, created_at)
			VALUES($1, $2, $3, $4, $5, $6, $7)`,
			id,
			rtransaction.StaffID,
			rtransaction.ProductID,
			rtransaction.RepositoryTransactionType,
			rtransaction.Price,
			rtransaction.Quantity,
			createdAt,
	); err != nil {
        log.Println("Error while inserting data:", err)
        return "", err
    }

    return id, nil
}

func (s *repositoryTransactionRepo) GetByID(id models.PrimaryKey) (models.RepositoryTransaction, error) {
	rtransaction := models.RepositoryTransaction{}
	query := `SELECT id, staff_id, product_id, repository_transaction_type, price, quantity, created_at, updated_at, deleted_at FROM repository_transactions WHERE id = $1`

	err := s.DB.QueryRow(context.Background(), query, id.ID).Scan(
		&rtransaction.ID,
		&rtransaction.StaffID,
		&rtransaction.ProductID,
		&rtransaction.RepositoryTransactionType,
		&rtransaction.Price,
		&rtransaction.Quantity,
		&rtransaction.UpdatedAt,
		&rtransaction.DeletedAt,
	)
	if err != nil {
		log.Println("Error while selecting repository by ID:", err)
		return models.RepositoryTransaction{}, err
	}

	return rtransaction, nil
}

func (s *repositoryTransactionRepo) GetList(req models.GetListRequest) (models.RepositoryTransactionsResponse, error) {
	var (
		rtransactions = []models.RepositoryTransaction{}
		count  int
	)

	countQuery := `SELECT COUNT(*) FROM repository_transactions`
	if req.Search != "" {
		countQuery += fmt.Sprintf(`WHERE quantity ILIKE '%%%s%%'`, req.Search)
	}

	err := s.DB.QueryRow(context.Background(), countQuery).Scan(&count)
	if err != nil {
		log.Println("Error while scanning count of repository_transactions:", err)
		return models.RepositoryTransactionsResponse{}, err
	}

	query := `SELECT id, staff_id, product_id, repository_transaction_type, price, quantity, created_at, updated_at, deleted_at FROM repository_transactions where deleted_at is null`
	if req.Search != "" {
		query += fmt.Sprintf(` WHERE quantity ILIKE '%%%s%%'`, req.Search)
	}
	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(context.Background(), query, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		log.Println("Error while querying repository_transactions:", err)
		return models.RepositoryTransactionsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		rtransaction := models.RepositoryTransaction{}
		err := rows.Scan(
			&rtransaction.ID,
			&rtransaction.StaffID,
			&rtransaction.ProductID,
			&rtransaction.RepositoryTransactionType,
			&rtransaction.Price,
			&rtransaction.Quantity,
			&rtransaction.UpdatedAt,
			&rtransaction.DeletedAt,
		)
		if err != nil {
			log.Println("Error while scanning row of repository_transactions:", err)
			return models.RepositoryTransactionsResponse{}, err
		}
		rtransactions = append(rtransactions, rtransaction)
	}

	return models.RepositoryTransactionsResponse{
		RepositoryTransactions: rtransactions,
		Count:  count,
	}, nil
}

func (s *repositoryTransactionRepo) Update(rtransaction models.UpdateRepositoryTransaction) (string, error) {
	query := `UPDATE repository_transactions SET staff_id = $1, product_id = $2, repository_transaction_type = $3, price = $4, quantity = $5, updated_at = NOW() WHERE id = $6`

	_, err := s.DB.Exec(context.Background(), query,
		&rtransaction.StaffID,
		&rtransaction.ProductID,
		&rtransaction.RepositoryTransactionType,
		&rtransaction.Price,
		&rtransaction.Quantity,
		&rtransaction.ID,
	)
	if err != nil {
		log.Println("Error while repository_transactions Repository :", err)
		return "", err
	}

	return rtransaction.ID, nil
}

func (s *repositoryTransactionRepo) Delete(id string) error {
	query := `UPDATE repository_transactions SET deleted_at = NOW() WHERE id = $1`

	_, err := s.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error while deleting repository_transactions ", err)
		return err
	}

	return nil
}