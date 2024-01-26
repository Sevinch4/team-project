package postgres

import (
	"database/sql"
	"fmt"
	"teamProject/config"
	"teamProject/storage"
)

type Store struct {
	DB *sql.DB
}

func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return Store{DB: db}, nil
}

func (s Store) Close() {
	s.DB.Close()
}

func (s Store) Branch() storage.IBranchStorage {
	return NewBranchRepo(s.DB)
}

func (s Store) Sale() storage.ISaleStorage {
	return NewSaleRepo(s.DB)
}

func (s Store) Transaction() storage.ITransactionStorage {
	return NewTransactionRepo(s.DB)
}
