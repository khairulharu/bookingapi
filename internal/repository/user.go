package repository

import (
	"context"

	"github.com/khairulharu/bookingapi/domain"
	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &repositoryUser{
		db: db,
	}
}

func (r repositoryUser) FindByPhone(ctx context.Context, phone string) (user domain.User, err error) {
	err = r.db.Debug().WithContext(ctx).Table("users").Where("phone=?", phone).First(&user).Error
	return
}

func (r repositoryUser) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	err = r.db.Debug().WithContext(ctx).Table("users").Where("id=?", id).First(&user).Error
	return
}

func (r repositoryUser) Insert(ctx context.Context, user *domain.User) error {
	err := r.db.Debug().WithContext(ctx).Table("users").Save(&user).Error
	return err
}
