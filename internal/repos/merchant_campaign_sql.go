package repos

import (
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
func (m *MerchantCampaignSQLRepo) ListMerchantCampaignRepo(page int, size int, filter map[string]interface{}) ([]*models.MerchantCampaignResponse, error) {
	var records []*models.MerchantCampaignResponse
	query := m.db
	//authorName, ok := filter["author"]
	//if ok {
	//	query = query.Where("author = ?", authorName)
	//}
	var err error
	if page != -1 && size != -1 {
		offset := (page - 1) * size
		//err = query.Order("id ASC").Limit(size).Offset(offset).Find(&records).Error
		if err := query.Table("merchant_campaigns").Select("*").
			Joins("JOIN merchants on merchant_campaigns.merchant_id = merchants.id").
			Joins("JOIN campaigns on merchant_campaigns.campaign_id = campaigns.id").
			Joins("JOIN providers on campaigns.provider_id = providers.id").
			Joins("JOIN rules on campaigns.rule_id = rules.id").
			Joins("JOIN items on rules.item_id = items.id").Find(&records).Order("id ASC").Limit(size).Offset(offset).Error; err != nil {
			return nil, err
		}
	} else {
		err = query.Table("merchant_campaigns").Order("id ASC").Find(&records).Error
	}
	return records, err
}
