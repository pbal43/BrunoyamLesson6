package inmemory

import (
	"BrunoyamLesson6/internal/domain/cars/errors"
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
)

func (storage *Storage) GetAllCars() ([]carDomain.Car, error) {
	if len(storage.cars) == 0 {
		return []carDomain.Car{}, errors.ErrCarsNotFound
	}
	var cars []carDomain.Car
	for _, car := range storage.cars {
		cars = append(cars, car)
	}
	return cars, nil
}

func (storage *Storage) GetCarByID(id string) (carDomain.Car, error) {
	car, ok := storage.cars[id]
	if !ok {
		return carDomain.Car{}, errors.ErrCarNotFound
	}
	return car, nil
}

func (storage *Storage) UpdateAvailable(cid string) error {
	car := storage.cars[cid]
	car.Available = !car.Available
	storage.cars[cid] = car
	return nil
}

func (storage *Storage) GetAvailableCars() ([]carDomain.Car, error) {
	if len(storage.cars) == 0 {
		return nil, errors.ErrCarsNotFound
	}
	var cars []carDomain.Car
	for _, car := range storage.cars {
		if car.Available {
			cars = append(cars, car)
		}
	}
	if len(cars) == 0 {
		return nil, errors.ErrAvailableCarsNotFound
	}
	return cars, nil
}

func (storage *Storage) AddCar(car carDomain.Car) error {
	for key, c := range storage.cars {
		if c.Label == car.Label && c.Model == car.Model && c.Year == car.Year {
			c.Count = c.Count + car.Count
			storage.cars[key] = c
			return nil
		}
	}
	storage.cars[car.CID] = car
	return nil
}
