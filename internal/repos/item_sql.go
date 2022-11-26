package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type ItemSQLRepo struct {
	db *gorm.DB
}

func NewItemSQLRepo(db *gorm.DB) IItemRepo {
	return &ItemSQLRepo{
		db: db,
	}
}

func (r ItemSQLRepo) ListItemByIDs(IDs []int64) ([]*models.Item, error) {
	var records []*models.Item
	err := r.db.Where("id in (?)", IDs).First(&records).Error
	return records, err
}

func (r ItemSQLRepo) GetItemByProviderID(providerID int) ([]*models.Item, error) {
	var records []*models.Item
	err := r.db.Where("provider_id = ?", providerID).First(&records).Error
	return records, err
}
