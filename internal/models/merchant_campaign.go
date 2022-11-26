package models

type MerchantCampaign struct {
	ID         int64 `json:"id"`
	MerchantID int64 `json:"merchant_id"`
	CampaignID int64 `json:"campaign_id"`
}
