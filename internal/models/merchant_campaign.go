package models

import "time"

type MerchantCampaign struct {
	ID         int64 `json:"id"`
	MerchantID int64 `json:"merchant_id"`
	CampaignID int64 `json:"campaign_id"`
}

type MerchantCampaignResponse struct {
	MerchantName     string    `json:"merchant_name"`
	CampaignName     string    `json:"campaign_name"`
	DSC              string    `json:"dsc"`
	CampaignImageURL string    `json:"campaign_image_url"`
	ProviderName     string    `json:"provider_name"`
	RequiredObNum    int64     `json:"required_ob_num"`
	RewardScore      int64     `json:"reward_score"`
	ItemName         string    `json:"item_name"`
	ExpiredAt        time.Time `json:"expired_at"`
}
