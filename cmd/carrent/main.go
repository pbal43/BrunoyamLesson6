package main

import (
	"BrunoyamLesson6/internal"
	"BrunoyamLesson6/internal/repository/inmemory"
	"BrunoyamLesson6/internal/server"
	"fmt"
)

func main() {

	// конфигураця приложения
	fmt.Println("RentApi is starting")
	cfg := internal.ReadConfig()
	// конфигураця и создание хранилища
	db := inmemory.NewInMemoryStorage()
	// конфигурация и запуск веб-сервера
	srv := server.NewServer(cfg, db)
	err := srv.Run()
	if err != nil {
		panic(err)
	}

}
