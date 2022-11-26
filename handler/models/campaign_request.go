package models

import (
	"time"
)

type CampaignRequest struct {
	ProviderID       int64      `json:"provider_id"`
	Name             string     `json:"name"`
	Rule             Rule       `json:"rule"`
	DSC              string     `json:"dsc"`
	CampaignImageURL string     `json:"campaign_image_url"`
	ExpiredAt        *time.Time `json:"expired_at"`
}

type Rule struct {
	ItemID        int64 `json:"item_id"`
	RequiredObNum int64 `json:"required_ob_num"`
}
