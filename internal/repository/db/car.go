package db

import (
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type carStorage struct {
	db *pgx.Conn
}

func (cs *carStorage) GetAllCars() ([]carDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := cs.db.Query(ctx, "SELECT * FROM cars")
	if err != nil {
		return nil, err
	}

	var cars []carDomain.Car
	for rows.Next() {
		var car carDomain.Car
		if err := rows.Scan(&car.CID, &car.Label, &car.Model, &car.Year, &car.Available); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (us *userStorage) GetCarByID(cid string) (carDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var car carDomain.Car
	err := us.db.QueryRow(ctx, "SELECT * FROM cars WHERE cid = $1", cid).
		Scan(&car.CID, &car.Label, &car.Model, &car.Year, &car.Available)
	if err != nil {
		return carDomain.Car{}, err
	}

	return car, nil
}

func (cs *carStorage) GetAvailableCars() ([]carDomain.Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := cs.db.Query(ctx, "SELECT * FROM cars WHERE available = true and count > 0")
	if err != nil {
		return nil, err
	}

	var cars []carDomain.Car
	for rows.Next() {
		var car carDomain.Car
		if err := rows.Scan(&car.CID, &car.Label, &car.Model, &car.Year, &car.Available); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (cs *carStorage) AddCar(car carDomain.Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cs.db.Exec(ctx, "INSERT into cars (cid, lable, model, year, available, count) VALUES ($1, $2, $3, $4, $5, $6)",
		car.CID, car.Label, car.Model, car.Year, car.Available, car.Count)
	if err != nil {
		return err
	}
	return nil
}

// TODO: добавить в аргументы новый статус available
func (cs *carStorage) UpdateAvailable(cid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cs.db.Exec(ctx, "UPDATE cars SET available = false WHERE cid = $1", cid)
	if err != nil {
		return err
	}
	return nil
}
