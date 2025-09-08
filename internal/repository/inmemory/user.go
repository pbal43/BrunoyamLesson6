package inmemory

import (
	userErrors "BrunoyamLesson6/internal/domain/users/errors"
	userDomain "BrunoyamLesson6/internal/domain/users/models"
	"github.com/google/uuid"
)

func (s *Storage) SaveUser(user userDomain.User) error {
	for _, userInMemory := range s.users {
		if user.Email == userInMemory.Email || user.Phone == userInMemory.Phone {
			return userErrors.ErrorUserIsAlreadyExist
		}
	}
	uid := uuid.New().String()
	_, ok := s.users[uid] // если uid есть в мапе, то генерируем новый id
	if ok {
		uid = uuid.New().String()
	}
	user.Uuid = uid
	s.users[user.Uuid] = user
	return nil
}

func (s *Storage) GetUser(userReq userDomain.UserRequest) (userDomain.User, error) {
	for _, userInMemory := range s.users {
		if userReq.Email == userInMemory.Email {
			return userInMemory, nil
		}
	}
	return userDomain.User{}, userErrors.ErrorUserNotExist
}
