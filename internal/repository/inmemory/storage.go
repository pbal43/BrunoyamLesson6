package inmemory

import (
	carDomain "BrunoyamLesson6/internal/domain/cars/models"
	userDomain "BrunoyamLesson6/internal/domain/users/models"
)

type Storage struct {
	users map[string]userDomain.User
	cars  map[string]carDomain.Car
}

func NewInMemoryStorage() *Storage {
	return &Storage{
		users: make(map[string]userDomain.User),
		cars:  make(map[string]carDomain.Car),
	}
}
