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

type basketRepo struct {
	DB *pgxpool.Pool
}

func NewBasketRepo(DB *pgxpool.Pool) storage.IBasketRepo {
	return &basketRepo{
		DB: DB,
	}
}

func (s *basketRepo) Create(basket models.CreateBasket) (string, error) {
	id := uuid.New().String()
	createdAT := time.Now()

	if _, err := s.DB.Exec(context.Background(), `INSERT INTO baskets 
		(id, staff_id, product_id, storage_trunsaction_type, price, quantity, created_at)
			VALUES($1, $2, $3, $4, $5, $6, $7)`,
		id,
		basket.SaleID,
		basket.ProductID,
		basket.Quantity,
		basket.Price,
		createdAT,
	); err != nil {
        log.Println("Error while inserting data:", err)
        return "", err
    }

    return id, nil
}

func (s *basketRepo) GetByID(id models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}
	query := `SELECT id, sale_id, product_id, quantity, price, created_at, updated_at, deleted_at FROM baskets WHERE id = $1`
	err := s.DB.QueryRow(context.Background(), query, id.ID).Scan(
		&basket.ID,
		&basket.SaleID,
		&basket.ProductID,
		&basket.Quantity,
		&basket.Price,
		&basket.CreatedAt,
		&basket.UpdatedAt,
		&basket.DeletedAt,
	)
	if err != nil {
		log.Println("Error while selecting basket by ID:", err)
		return models.Basket{}, err
	}
	return basket, nil
}

func (s *basketRepo) GetList(request models.GetListRequest) (models.BasketsResponse, error) {
	var (
		baskets = []models.Basket{}
		count       int
	)

	countQuery := `SELECT COUNT(*) FROM baskets`
	if request.Search != "" {
		countQuery += fmt.Sprintf(` WHERE branch_id ILIKE '%%%s%%'`, request.Search)
	}

	err := s.DB.QueryRow(context.Background(), countQuery).Scan(&count)
	if err != nil {
		log.Println("Error while scanning count of repositories:", err)
		return models.BasketsResponse{}, err
	}

	query := `SELECT id, sale_id, product_id, quantity, price, created_at, updated_at, deleted_at FROM baskets where deleted_at is null`
	if request.Search != "" {
		query += fmt.Sprintf(` WHERE branch_id ILIKE '%%%s%%'`, request.Search)
	}
	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(context.Background(), query, request.Limit, (request.Page-1)*request.Limit)
	if err != nil {
		log.Println("Error while querying baskets:", err)
		return models.BasketsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		basket := models.Basket{}
		err := rows.Scan(
			&basket.ID,
			&basket.SaleID,
			&basket.ProductID,
			&basket.Quantity,
			&basket.Price,
			&basket.CreatedAt,
			&basket.UpdatedAt,
			&basket.DeletedAt,
		)
		if err != nil {
			log.Println("Error while scanning row of baskets:", err)
			return models.BasketsResponse{}, err
		}
		baskets = append(baskets, basket)
	}

	return models.BasketsResponse{
		Baskets: baskets,
		Count:       count,
	}, nil
}

func (s *basketRepo) Update(repository models.UpdateRepository) (string, error) {
	query := `UPDATE repositories SET branch_id = $1, product_id = $2, count = $3 WHERE id = $4`

	_, err := s.DB.Exec(context.Background(), query,
		&repository.BranchID,
		&repository.ProductID,
		&repository.Count,
		&repository.ID,
	)
	if err != nil {
		log.Println("Error while updating Repository :", err)
		return "", err
	}

	return repository.ID, nil
}