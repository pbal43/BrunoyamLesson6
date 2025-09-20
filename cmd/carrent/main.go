package main

import (
	"BrunoyamLesson6/internal"
	newdb "BrunoyamLesson6/internal/repository/db"
	"BrunoyamLesson6/internal/server"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {

	// конфигураця приложения
	fmt.Println("RentApi is starting")
	cfg := internal.ReadConfig()

	// конфигураця и создание хранилища
	//db := inmemory.NewInMemoryStorage()
	database, err := newdb.NewStorage(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// запуск миграции
	if err := newdb.Migrations(cfg.DSN, cfg.MigratePath); err != nil {
		log.Fatal(err)
	}

	// конфигурация и запуск веб-сервера
	srv := server.NewServer(cfg, database)
	if err := srv.Run(); err != nil {
		panic(err)
	}

}
