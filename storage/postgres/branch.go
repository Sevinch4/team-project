package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"teamProject/api/models"
	"teamProject/storage"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) storage.IBranchStorage {
	return branchRepo{db: db}
}
func (b branchRepo) Create(branch models.CreateBranch) (string, error) {
	id := uuid.New()

	query := `insert into branches (id, name, address) 
									values($1, $2, $3)`

	if _, err := b.db.Exec(query,
		id,
		branch.Name,
		branch.Address); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (b branchRepo) GetByID(id string) (models.Branch, error) {
	branch := models.Branch{}
	query := `select id, name, address, created_at, updated_at from branches where id = $1 and deleted_at is null`
	if err := b.db.QueryRow(query, id).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.CreatedAt,
		&branch.UpdatedAt); err != nil {
		fmt.Println("error is while selecting by id", err.Error())
		return models.Branch{}, err
	}
	return branch, nil
}

func (b branchRepo) GetList(request models.GetListRequest) (models.BranchResponse, error) {
	var (
		count             = 0
		branches          = []models.Branch{}
		query, countQuery string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `select count(1) from branches where deleted_at is NULL `

	if search != "" {
		countQuery += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}

	if err := b.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count", err.Error())
		return models.BranchResponse{}, err
	}

	query = `select id, name, address, created_at, updated_at from branches where deleted_at is NULL  `
	if search != "" {
		query += fmt.Sprintf(` and name ilike '%%%s%%' `, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting * from branches", err.Error())
		return models.BranchResponse{}, err
	}

	for rows.Next() {
		branch := models.Branch{}
		if err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.CreatedAt,
			&branch.UpdatedAt); err != nil {
			fmt.Println("error is while scanning branch", err.Error())
			return models.BranchResponse{}, err
		}
		branches = append(branches, branch)
	}

	return models.BranchResponse{
		Branches: branches,
		Count:    count,
	}, err
}
func (b branchRepo) Update(branch models.UpdateBranch) (string, error) {
	query := `update branches set name = $1, address = $2, updated_at = Now() where id = $3`

	if _, err := b.db.Exec(query,
		&branch.Name,
		&branch.Address,
		&branch.ID); err != nil {
		fmt.Println("error is while updating branch", err.Error())
		return "", err
	}

	return branch.ID, nil
}
func (b branchRepo) Delete(id string) error {
	query := `update branches set deleted_at = now() where id = $1`

	if _, err := b.db.Exec(query, id); err != nil {
		fmt.Println("error is while deleting branches", err.Error())
		return err
	}
	return nil
}
