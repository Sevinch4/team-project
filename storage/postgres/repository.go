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

type repositoryRepo struct {
	DB *pgxpool.Pool
}

func NewRepositoryRepo(DB *pgxpool.Pool) storage.IRepositoryRepo {
	return &repositoryRepo{
		DB: DB,
	}
}

func (s *repositoryRepo) Create(repository models.CreateRepository) (string, error) {
    id := uuid.New().String()
    createdAt := time.Now()

    if _, err := s.DB.Exec(context.Background(), `INSERT INTO repositories 
    (id, product_id, branch_id, count, created_at) 
        VALUES ($1, $2, $3, $4, $5)`,
        id,
        repository.ProductID,
        repository.BranchID,
        repository.Count,
        createdAt,
    ); err != nil {
        log.Println("Error while inserting data:", err)
        return "", err
    }

    return id, nil
}

func (s *repositoryRepo) GetByID(id models.PrimaryKey) (models.Repository, error) {
	repository := models.Repository{}
	query := `SELECT id, product_id, branch_id, count, created_at, updated_at, deleted_at FROM repositories WHERE id = $1`
	err := s.DB.QueryRow(context.Background(), query, id.ID).Scan(
		&repository.ID,
		&repository.ProductID,
		&repository.BranchID,
		&repository.Count,
		&repository.CreatedAt,
		&repository.UpdatedAt,
		&repository.DeletedAt,
	)
	if err != nil {
		log.Println("Error while selecting repository by ID:", err)
		return models.Repository{}, err
	}
	return repository, nil
}

func (s *repositoryRepo) GetList(request models.GetListRequest) (models.RepositoriesResponse, error) {
	var (
		repositories = []models.Repository{}
		count       int
	)

	countQuery := `SELECT COUNT(*) FROM repositories`
	if request.Search != "" {
		countQuery += fmt.Sprintf(` WHERE branch_id ILIKE '%%%s%%'`, request.Search)
	}

	err := s.DB.QueryRow(context.Background(), countQuery).Scan(&count)
	if err != nil {
		log.Println("Error while scanning count of repositories:", err)
		return models.RepositoriesResponse{}, err
	}

	query := `SELECT id, product_id, branch_id, count, created_at, updated_at, deleted_at FROM repositories where deleted_at is null`
	if request.Search != "" {
		query += fmt.Sprintf(` WHERE branch_id ILIKE '%%%s%%'`, request.Search)
	}
	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(context.Background(), query, request.Limit, (request.Page-1)*request.Limit)
	if err != nil {
		log.Println("Error while querying repositories:", err)
		return models.RepositoriesResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		repository := models.Repository{}
		err := rows.Scan(
			&repository.ID,
			&repository.ProductID,
			&repository.BranchID,
			&repository.Count,
			&repository.CreatedAt,
			&repository.UpdatedAt,
			&repository.DeletedAt,
		)
		if err != nil {
			log.Println("Error while scanning row of repositories:", err)
			return models.RepositoriesResponse{}, err
		}
		repositories = append(repositories, repository)
	}

	return models.RepositoriesResponse{
		Repositories: repositories,
		Count:       count,
	}, nil
}

func (s *repositoryRepo) Update(repository models.UpdateRepository) (string, error) {
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

func (s *repositoryRepo) Delete(id string) error {
	query := `UPDATE repositories SET deleted_at = NOW() WHERE id = $1`

	_, err := s.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error while deleting Repository :", err)
		return err
	}

	return nil
}
