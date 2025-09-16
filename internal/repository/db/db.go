package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	userStorage
	carStorage
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &Storage{
		userStorage: userStorage{db},
		carStorage:  carStorage{db},
	}, nil
}
