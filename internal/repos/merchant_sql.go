package repos

import (
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

func (r MerchantSQLRepo) ListMerchantByIDs(IDs []int64) ([]*models.Merchant, error) {
	var records []*models.Merchant
	err := r.db.Where("id in (?)", IDs).First(&records).Error
	return records, err
}

func (r MerchantSQLRepo) Create(record *models.Merchant) (*models.Merchant, error) {
	err := r.db.Create(record).Error
	return record, err
}
