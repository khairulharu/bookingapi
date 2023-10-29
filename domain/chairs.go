package domain

import "context"

type Chair struct {
	Id        int64  `db:"id"`
	IsBook    int8   `db:"is_book"`
	UserBook  string `db:"user_book"`
	UserPhone string `db:"user_phone"`
	Pay       int8   `db:"pay_status"`
}

type ChairRepository interface {
	GetChairs(ctx context.Context) ([]Chair, error)
	GetChairByID(ctx context.Context, id int64) (Chair, error)
	UpdateChair(ctx context.Context, chair *Chair) error
}
