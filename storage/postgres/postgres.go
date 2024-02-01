package postgres

import (
	"context"
	"fmt"
	"teamProject/config"
	"teamProject/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	poolConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB))
	if err != nil {
		fmt.Println("error is while parsing config", err.Error())
		return nil, err
	}
	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		fmt.Println("error is while connecting to db", err.Error())
		return nil, err
	}
	return &Store{
		Pool: pool,
		}, nil
}


func (s *Store) Close() {
	s.Pool.Close()
}

func (s *Store) StaffTarif() storage.IStaffTarifRepo {
	return NewStaffTarifRepo(s.Pool)
}

func (s *Store) Staff() storage.IStaffRepo {
	return NewStaffRepo(s.Pool)
}

func (s *Store) Repository() storage.IRepositoryRepo {
	return NewRepositoryRepo(s.Pool)
}

func (s *Store) Basket() storage.IBasketRepo {
	return 	NewBasketRepo(s.Pool)
}

func (s *Store) RTransaction() storage.IRepositoryTransactionRepo {
	return NewRepositoryTransactionRepo(s.Pool)
}
