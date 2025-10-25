package db

import (
	"BrunoyamLesson6/internal/domain"
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type carStorage struct {
	db *pgx.Conn
}

func (cs *carStorage) GetAllCars() ([]carDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	rows, err := cs.db.Query(ctx, "SELECT cid, lable, model, year, available, price FROM cars")
	if err != nil {
		return nil, err
	}

	var cars []carDomain.Car
	for rows.Next() {
		var car carDomain.Car
		if err = rows.Scan(&car.CID, &car.Label, &car.Model, &car.Year, &car.Available); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (us *userStorage) GetCarByID(cid string) (carDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	var car carDomain.Car
	err := us.db.QueryRow(ctx, "SELECT cid, lable, model, year, available, price FROM cars WHERE cid = $1", cid).
		Scan(&car.CID, &car.Label, &car.Model, &car.Year, &car.Available)
	if err != nil {
		return carDomain.Car{}, err
	}

	return car, nil
}

func (cs *carStorage) GetAvailableCars() ([]carDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	rows, err := cs.db.Query(ctx, "SELECT cid, lable, model, year, available, price FROM cars WHERE available = true")
	if err != nil {
		return nil, err
	}

	var cars []carDomain.Car
	for rows.Next() {
		var car carDomain.Car
		if err = rows.Scan(&car.CID, &car.Label, &car.Model, &car.Year, &car.Available); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (cs *carStorage) AddCar(car carDomain.Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	_, err := cs.db.Exec(
		ctx,
		"INSERT into cars (cid, lable, model, year, available) VALUES ($1, $2, $3, $4, $5)",
		car.CID,
		car.Label,
		car.Model,
		car.Year,
		car.Available,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cs *carStorage) AddCars(cars []carDomain.Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := cs.db.Begin(ctx) // открыли транзакцию
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // если упало - откатываем изменения, но если ниже будет коммит - то зафиксируем изменения и не откатит

	_, err = tx.Prepare(ctx, "add_car", "INSERT into cars (cid, lable, model, year, available) VALUES ($1, $2, $3, $4, $5)") // подготовленный запрос
	if err != nil {
		return err
	}

	for _, car := range cars { // засунули проход по всем в 1 транзакцию
		_, err = tx.Exec(ctx, "add_car", car.CID, car.Label, car.Model, car.Year, car.Available)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx) // коммитим
}

// TODO: добавить в аргументы новый статус available
func (cs *carStorage) UpdateAvailable(cid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	_, err := cs.db.Exec(ctx, "UPDATE cars SET available = false WHERE cid = $1", cid)
	if err != nil {
		return err
	}
	return nil
}
