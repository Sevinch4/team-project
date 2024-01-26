package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"teamProject/config"
	"teamProject/storage"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (*Store, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
	}, nil
}

func (s *Store) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

func (s *Store) StaffTarif() storage.IStaffTarifRepo {
	return NewStaffTarifRepo(s.DB)
}

func (s *Store) Staff() storage.IStaffRepo {
	return NewStaffRepo(s.DB)
}
