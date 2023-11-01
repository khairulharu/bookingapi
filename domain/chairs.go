package domain

import (
	"context"

	"github.com/khairulharu/bookingapi/dto"
)

type Chair struct {
	Id        int64  `db:"id"`
	CodeRef   string `db:"code_ref"`
	IsBook    int8   `db:"is_book"`
	UserBook  string `db:"user_book"`
	UserPhone string `db:"user_phone"`
	Pay       int8   `db:"pay"`
}

type ChairRepository interface {
	Insert(ctx context.Context, chair *[]Chair, value int) error
	GetChairs(ctx context.Context) ([]Chair, error)
	GetChairByID(ctx context.Context, id int64) (Chair, error)
	Update(ctx context.Context, chair *Chair) error
	Delete(ctx context.Context, chairs *[]Chair) error
}

type ChairService interface {
	StoreChairs(ctx context.Context, value int) dto.Response
	DeleteChairs(ctx context.Context) dto.Response
}
