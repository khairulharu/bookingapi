package domain

import "context"

type User struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Phone string `db:"phone"`
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
}
