package domain

import (
	"context"

	"github.com/khairulharu/bookingapi/dto"
)

type User struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	Delete(ctx context.Context, users *[]User) error
}

type UserService interface {
	StoreUser(ctx context.Context, req dto.ReqUser) dto.Response
}
