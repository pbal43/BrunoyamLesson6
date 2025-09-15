package carusecase

import (
	"BrunoyamLesson6/internal/domain/cars/errors"
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
	"github.com/google/uuid"
)

type CarStorage interface {
	GetAllCars() ([]carDomain.Car, error)
	GetCarByID(id string) (carDomain.Car, error)
	GetAvailableCars() ([]carDomain.Car, error)
	AddCar(car carDomain.Car) error
	UpdateAvailable(cid string) error
}

type CarUsecase struct {
	db CarStorage
}

func NewCarUsecase(db CarStorage) *CarUsecase {
	return &CarUsecase{db: db}
}

func (c CarUsecase) GetAllCars() ([]carDomain.Car, error) {
	return c.db.GetAllCars()
}

func (c CarUsecase) GetCarByID(id string) (carDomain.Car, error) {
	car, err := c.db.GetCarByID(id)
	if err != nil {
		return carDomain.Car{}, err
	}
	if !car.Available {
		return carDomain.Car{}, errors.ErrCarNotAvailable
	}
	car.Available = false
	if err := c.db.UpdateAvailable(id); err != nil {
		return carDomain.Car{}, err
	}
	return car, nil
}

func (c CarUsecase) AddCar(car carDomain.Car) error {
	cid := uuid.New().String()
	car.CID = cid
	if car.Count == 0 {
		car.Count = 1
	}
	err := c.db.AddCar(car)
	if err != nil {
		return err
	}
	return nil
}

func (c CarUsecase) GetAvailableCars() ([]carDomain.Car, error) {
	return c.db.GetAvailableCars()
}
