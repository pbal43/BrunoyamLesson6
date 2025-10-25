package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	userStorage
	carStorage
}

func (s *Storage) Close(ctx context.Context) error {
	return s.userStorage.db.Close(ctx)
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

func Migrations(dsn string, migratePath string) error {

	mPath := fmt.Sprintf("file://%s", migratePath)
	m, err := migrate.New(mPath, dsn)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		log.Println("DB is already up to date")
	}

	log.Println("Migration complete")

	return nil
}
