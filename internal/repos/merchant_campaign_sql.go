package repos

import (
	"context"
	handlerModels "github.com/vietpham1023/golang-uit-hackathon/handler/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type MerchantCampaignSQLRepo struct {
	db *gorm.DB
}

func NewMerchantCampaignSQLRepo(db *gorm.DB) IMerchantCampaignRepo {
	return &MerchantCampaignSQLRepo{
		db: db,
	}
}
func (m *MerchantCampaignSQLRepo) ListMerchantCampaignRepo(page int, size int, filter map[string]interface{}) ([]*models.MerchantCampaign, error) {
	var records []*models.MerchantCampaign
	query := m.db
	//authorName, ok := filter["author"]
	//if ok {
	//	query = query.Where("author = ?", authorName)
	//}
	var err error
	if page != -1 && size != -1 {
		offset := (page - 1) * size
		err = query.
			Order("created_at DESC").
			Limit(size).
			Offset(offset).
			Find(&records).Error
	} else {
		err = query.Order("created_at DESC").Find(&records).Error
	}
	return records, err
}

func (m *MerchantCampaignSQLRepo) Create(ctx context.Context, record *handlerModels.MerchantCampaign) error {
	err := m.db.Create(record).Error
	if err != nil {
		return err
	}
	return nil
}
