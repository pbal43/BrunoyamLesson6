package main

import (
	"BrunoyamLesson6/internal"
	newdb "BrunoyamLesson6/internal/repository/db"
	"BrunoyamLesson6/internal/server"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		defer signal.Stop(c)
		<-c
		cancel()
	}()

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
	srv := server.NewServer(cfg, database)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// конфигурация и запуск веб-сервера
		if err = srv.Run(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		err = srv.ShutDown(ctx)
		if err != nil {
			log.Fatal(err)
		}
		err = database.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("To-do-list Api is shutting down")
	}()
	wg.Wait()
}
