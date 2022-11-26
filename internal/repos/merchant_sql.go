package repos

import (

	"context"
	handlerModels "github.com/vietpham1023/golang-uit-hackathon/handler/models"

	"github.com/vietpham1023/golang-uit-hackathon/internal/models"

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

func (r MerchantSQLRepo) ListMerchantByIDs(IDs []int64) ([]*models.Merchant, error) {
	var records []*models.Merchant
	err := r.db.Where("id in (?)", IDs).First(&records).Error
	return records, err
}

func (r MerchantSQLRepo) Create(record *models.Merchant) (*models.Merchant, error) {
	err := r.db.Create(record).Error
	return record, err
}
