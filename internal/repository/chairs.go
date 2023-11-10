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

func (r repositoryChair) Insert(ctx context.Context, chair *[]domain.Chair, value int) error {
	err := r.db.Debug().WithContext(ctx).Table("chairs").CreateInBatches(chair, value).Error
	return err
}

func (r repositoryChair) GetChairs(ctx context.Context) (chairs []domain.Chair, err error) {
	err = r.db.Debug().WithContext(ctx).Table("chairs").Order("id").Find(&chairs).Error
	return
}

func (r repositoryChair) GetChairByID(ctx context.Context, id int64) (chair domain.Chair, err error) {
	err = r.db.Debug().WithContext(ctx).Table("chairs").Where("id=?", id).First(&chair).Error
	return
}

func (r repositoryChair) GetChairByCode(ctx context.Context, code string) (chair domain.Chair, err error) {
	err = r.db.Debug().WithContext(ctx).Table("chairs").Where("code_ref=?", code).First(&chair).Error
	return
}
func (r repositoryChair) Update(ctx context.Context, chair *domain.Chair) error {
	err := r.db.Debug().WithContext(ctx).Table("chairs").Save(&chair).Error
	return err
}

func (r repositoryChair) Delete(ctx context.Context, chairs *[]domain.Chair) error {
	err := r.db.Debug().WithContext(ctx).Table("chairs").Delete(*chairs).Error
	return err
}
