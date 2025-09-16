package main

import (
	"BrunoyamLesson6/internal"
	newdb "BrunoyamLesson6/internal/repository/db"
	"BrunoyamLesson6/internal/server"
	"fmt"
	"log"
)

func main() {

	// конфигураця приложения
	fmt.Println("RentApi is starting")
	cfg := internal.ReadConfig()
	// конфигураця и создание хранилища
	//db := inmemory.NewInMemoryStorage()
	db, err := newdb.NewStorage(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}
	// конфигурация и запуск веб-сервера
	srv := server.NewServer(cfg, db)
	if err := srv.Run(); err != nil {
		panic(err)
	}

}
