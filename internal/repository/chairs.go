package repository

import (
	"context"

	"github.com/khairulharu/bookingapi/domain"
	"gorm.io/gorm"
)

type repositoryChair struct {
	db *gorm.DB
}

func NewChair(db *gorm.DB) domain.ChairRepository {
	return &repositoryChair{
		db: db,
	}
}

func (r repositoryChair) GetChairs(ctx context.Context) (chairs []domain.Chair, err error) {
	err = r.db.Debug().WithContext(ctx).Table("chairs").Order("id").Find(&chairs).Error
	return
}

func (r repositoryChair) GetChairByID(ctx context.Context, id int64) (chair domain.Chair, err error) {
	err = r.db.Debug().WithContext(ctx).Table("chairs").Where("id=?", id).First(&chair).Error
	return
}

func (r repositoryChair) UpdateChair(ctx context.Context, chair *domain.Chair) error {
	err := r.db.Debug().WithContext(ctx).Table("chairs").Save(&chair).Error
	return err
}
