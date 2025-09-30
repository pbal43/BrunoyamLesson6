package main

import (
	"BrunoyamLesson6/internal"
	newdb "BrunoyamLesson6/internal/repository/db"
	"BrunoyamLesson6/internal/server"
	"BrunoyamLesson6/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// конфигураця приложения
	cfg := internal.ReadConfig()

	// указание уровня логирования
	log := logger.Init(cfg.Debug)
	log.Debug().Any("config", cfg).Send()
	log.Info().Msg("RentAPI is starting")

	// конфигураця и создание хранилища
	// db := inmemory.NewInMemoryStorage()
	database, err := newdb.NewStorage(cfg.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// запуск миграции
	if err = newdb.Migrations(cfg.DSN, cfg.MigratePath, &log); err != nil {
		log.Fatal().Err(err).Msg("Failed to create migrations")
	}

	// конфигурация и запуск веб-сервера
	srv := server.NewServer(cfg, database, &log)
	if err = srv.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}
}
