package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type ProviderSQLRepo struct {
	db *gorm.DB
}

func NewProviderSQLRepo(db *gorm.DB) IProviderRepo {
	return &ProviderSQLRepo{
		db: db,
	}
}

func (p ProviderSQLRepo) ListProviderByIDs(IDs []int64) ([]*models.Provider, error) {
	var records []*models.Provider
	err := p.db.Where("id in (?)", IDs).First(&records).Error
	return records, err
}

func (p ProviderSQLRepo) Create(record *models.Provider) (*models.Provider, error) {
	err := p.db.Create(record).Error
	return record, err
}
