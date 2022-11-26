package repos

import (
	"context"
	handlerModels "github.com/vietpham1023/golang-uit-hackathon/handler/models"
	"gorm.io/gorm"
)

type MerchantSQLRepo struct {
	db *gorm.DB
}

func NewMerchantSQLRepo(db *gorm.DB) IMerchantRepo {
	return &MerchantSQLRepo{
		db: db,
	}
}

func (m MerchantSQLRepo) CreateMerchant(ctx context.Context, merchant *handlerModels.Merchant) error {
	err := m.db.Create(merchant).Error
	return err
}

func (m MerchantSQLRepo) FindByID(ctx context.Context, id int) (*handlerModels.Merchant, error) {
	var merchant handlerModels.Merchant
	err := m.db.Where("id = ?", id).First(&merchant).Error
	return &merchant, err
}

func (m MerchantSQLRepo) ListByConditions(ctx context.Context, filter map[string]interface{}) ([]*handlerModels.Merchant, error) {
	var merchant []*handlerModels.Merchant
	err := m.db.Table("merchants").Where(filter).First(&merchant).Error
	return merchant, err
}
