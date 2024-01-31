package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"teamProject/api/models"
	"teamProject/storage"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.IProducts {
	return productRepo{db: db}
}

func (p productRepo) Create(product models.CreateProduct) (string, error) {
	id := uuid.New()
	query := `insert into products (id, name, price, barcode, category_id) values($1, $2, $3, $4, $5)`
	if _, err := p.db.Exec(context.Background(), query,
		id, product.Name, product.Price, product.Barcode, product.CategoryID); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return "", err
	}
	return id.String(), nil
}

func (p productRepo) GetByID(id string) (models.Product, error) {
	product := models.Product{}
	query := `select id, name, price, barcode, category_id, created_at, updated_at 
							from products where id = $1 and deleted_at is null`
	if err := p.db.QueryRow(context.Background(), query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt); err != nil {
		fmt.Println("error is while scanning", err.Error())
		return models.Product{}, err
	}
	return product, nil
}

func (p productRepo) GetList(request models.GetListRequest) (models.ProductResponse, error) {
	var (
		page              = request.Page
		offset            = (page - 1) * request.Limit
		query, countQuery string
		count             = 0
		products          = []models.Product{}
		search            = request.Search
	)
	countQuery = `select count(1) from products where deleted_at is null `
	if search != "" {
		countQuery += fmt.Sprintf(` and (name ilike '%%%s%%' or CAST(barcode as text) ilike '%%%s%%')`, search, search)
	}
	if err := p.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count", err.Error())
		return models.ProductResponse{}, err
	}

	query = `select  id, name, price, barcode, category_id, created_at, updated_at 
							from products where deleted_at is null `
	if search != "" {
		query += fmt.Sprintf(` and (name ilike '%%%s%%' or CAST(barcode as text) ilike '%%%s%%') `, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`
	rows, err := p.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting all", err.Error())
		return models.ProductResponse{}, err
	}

	for rows.Next() {
		product := models.Product{}
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt); err != nil {
			fmt.Println("error is while scanning category", err.Error())
			return models.ProductResponse{}, err
		}
		products = append(products, product)
	}
	return models.ProductResponse{
		Products: products,
		Count:    count,
	}, nil

}

func (p productRepo) Update(product models.UpdateProduct) (string, error) {
	query := `update products set name = $1, price = $2, category_id = $3, updated_at = now() 
									where id = $4`
	if _, err := p.db.Exec(context.Background(), query,
		&product.Name,
		&product.Price,
		&product.CategoryID,
		&product.ID); err != nil {
		fmt.Println("error is while updating", err.Error())
		return "", err
	}
	return product.ID, nil
}

func (p productRepo) Delete(id string) error {
	query := `update products set deleted_at = now() where id = $1`
	if _, err := p.db.Exec(context.Background(), query, &id); err != nil {
		fmt.Println("error is while deleting", err.Error())
		return err
	}
	return nil
}
