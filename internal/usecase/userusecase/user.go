package userusecase

import (
	"BrunoyamLesson6/internal/domain/users/errors"
	"BrunoyamLesson6/internal/domain/users/models"
	"github.com/go-playground/validator/v10"
)

type UserStorage interface {
	SaveUser(user models.User) error
	GetUser(userReq models.UserRequest) (models.User, error)
}

type UserUsecase struct {
	db    UserStorage
	valid *validator.Validate
}

func NewUserUsecase(db UserStorage) *UserUsecase {
	return &UserUsecase{db: db, valid: validator.New()}
}

func (uu *UserUsecase) SaveUser(user models.User) error {
	err := uu.valid.Struct(user)
	if err != nil {
		return err
	}
	// TODO: Генерация хеша
	return uu.db.SaveUser(user)
}

func (uu *UserUsecase) LoginUser(userReq models.UserRequest) (models.User, error) {
	dbUser, err := uu.db.GetUser(userReq)
	if err != nil {
		return models.User{}, err
	}

	// TODO: Проверка хеша пароля
	if dbUser.Password != userReq.Password {
		return models.User{}, errors.ErrorInvalidPassword
	}
	return dbUser, nil
}
