package models

type MerchantCampaign struct {
	ID         int64   `json:"id" gorm:"id"`
	MerchantID int64   `json:"merchant_id" gorm:"merchant_id"`
	CampaignID int64   `json:"campaign_id" gorm:"campaign_id"`
	Lat        float32 `json:"lat" gorm:"lat"`
	Lng        float32 `json:"lng" gorm:"long"`
}

type ListMerchantCampaign struct {
	MerchantID int64 `url:"merchant_id" gorm:"merchant_id"`
	CampaignID int64 `url:"campaign_id" gorm:"campaign_id"`
}
