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

type staffRepo struct {
	DB *pgxpool.Pool
}

func NewStaffRepo(DB *pgxpool.Pool) storage.IStaffRepo {
	return &staffRepo{
		DB: DB,
	}
}

func (s *staffRepo) Create(staff models.CreateStaff) (string, error) {
	id := uuid.New().String()
	createdAt := time.Now()

	birthDate, err := time.Parse("2006-01-02", staff.BirthDate)
	if err != nil {
		log.Println("Error parsing birth date:", err)
		return "", err
	}
	age := uint(time.Since(birthDate).Hours() / 24 / 365)

	if _, err := s.DB.Exec(context.Background(), `INSERT INTO staffs 
		(id, branch_id, tariff_id, staff_type, name, balance, age, birth_date, login, password, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		id,
		staff.BranchID,
		staff.TariffID,
		staff.StaffType,
		staff.Name,
		staff.Balance,
		age,
		birthDate,
		staff.Login,
		staff.Password,
		createdAt,
	); err != nil {
		log.Println("Error while inserting data ", err)
		return "", err
	}
	return id, nil
}

func (s *staffRepo) StaffByID(id models.PrimaryKey) (models.Staff, error) {
	staff := models.Staff{}
	query := `SELECT id, branch_id, tariff_id, staff_type, name, balance, age, birth_date, login, created_at, updated_at FROM staffs WHERE id = $1`

	err := s.DB.QueryRow(context.Background(), query, id.ID).Scan(
		&staff.ID,
		&staff.BranchID,
		&staff.TariffID,
		&staff.StaffType,
		&staff.Name,
		&staff.Balance,
		&staff.Age,
		&staff.BirthDate,
		&staff.Login,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	)
	if err != nil {
		log.Println("Error while selecting staff by ID:", err)
		return models.Staff{}, err
	}

	return staff, nil
}

func (s *staffRepo) GetStaffTList(request models.GetListRequest) (models.StaffsResponse, error) {
	var (
		staffs = []models.Staff{}
		count  int
	)

	countQuery := `SELECT COUNT(*) FROM staffs`
	if request.Search != "" {
		countQuery += fmt.Sprintf(` WHERE name ILIKE '%%%s%%'`, request.Search)
	}

	err := s.DB.QueryRow(context.Background(), countQuery).Scan(&count)
	if err != nil {
		log.Println("Error while scanning count of staffs:", err)
		return models.StaffsResponse{}, err
	}

	query := `SELECT id, branch_id, tariff_id, staff_type, name, balance, age, birth_date, login, created_at, updated_at FROM staffs where deleted_at is null`
	if request.Search != "" {
		query += fmt.Sprintf(` WHERE name ILIKE '%%%s%%'`, request.Search)
	}
	query += ` LIMIT $1 OFFSET $2`

	rows, err := s.DB.Query(context.Background(), query, request.Limit, (request.Page-1)*request.Limit)
	if err != nil {
		log.Println("Error while querying staff :", err)
		return models.StaffsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		staff := models.Staff{}
		err := rows.Scan(
			&staff.ID,
			&staff.BranchID,
			&staff.TariffID,
			&staff.StaffType,
			&staff.Name,
			&staff.Balance,
			&staff.Age,
			&staff.BirthDate,
			&staff.Login,
			&staff.CreatedAt,
			&staff.UpdatedAt,
		)
		if err != nil {
			log.Println("Error while scanning row of staffs:", err)
			return models.StaffsResponse{}, err
		}
		staffs = append(staffs, staff)
	}

	return models.StaffsResponse{
		Staffs: staffs,
		Count:  count,
	}, nil
}

func (s *staffRepo) UpdateStaff(staff models.UpdateStaff) (string, error) {
	query := `UPDATE staffs SET branch_id = $1, tariff_id = $2, staff_type = $3, name = $4, balance = $5, login = $6, updated_at = NOW() WHERE id = $7`

	_, err := s.DB.Exec(context.Background(), query,
		&staff.BranchID,
		&staff.TariffID,
		&staff.StaffType,
		&staff.Name,
		&staff.Balance,
		&staff.Login,
		staff.ID,
	)
	if err != nil {
		log.Println("Error while updating Staff :", err)
		return "", err
	}

	return staff.ID, nil
}

func (s *staffRepo) DeleteStaff(id string) error {
	query := `UPDATE staffs SET deleted_at = NOW() WHERE id = $1`

	_, err := s.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error while deleting Staff :", err)
		return err
	}

	return nil
}

func (s *staffRepo) GetPassword(id string) (string, error) {
	password := ""

	query := `
		select password from staffs 
		                where  id = $1`

	if err := s.DB.QueryRow(context.Background(), query, id).Scan(&password); err != nil {
		fmt.Println("Error while scanning password from users", err.Error())
		return "", err
	}

	return password, nil
}

func (s *staffRepo) UpdatePassword(request models.UpdateStaffPassword) error {
	query := `
		update staffs 
				set password = $1
					where id = $2`

	if _, err := s.DB.Exec(context.Background(), query, request.NewPassword, request.ID); err != nil {
		fmt.Println("error while updating password for staff", err.Error())
		return err
	}

	return nil
}
