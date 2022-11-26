package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type CampaignSQLRepo struct {
	db *gorm.DB
}

func NewCampaignSQLRepo(db *gorm.DB) ICampaignRepo {
	return &CampaignSQLRepo{
		db: db,
	}
}

func (c CampaignSQLRepo) ListCampaignByIDs(IDs []int64) ([]*models.Campaign, error) {
	var records []*models.Campaign
	err := c.db.Where("id in (?)", IDs).Order("").Find(&records).Error
	return records, err
}

func (c CampaignSQLRepo) Create(record *models.Campaign) (*models.Campaign, error) {
	err := c.db.Create(record).Error
	return record, err
}
