package models

type User struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Age      int    `json:"age" validate:"required,gte=18"`
	Phone    string `json:"phone" validate:"required,e164"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
