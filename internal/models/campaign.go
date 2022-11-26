package models

import "time"

type Campaign struct {
	ID               int64     `json:"id"`
	ProviderID       int64     `json:"provider_id"`
	Name             string    `json:"name"`
	DSC              string    `json:"dsc"` // description
	RuleID           int64     `json:"RuleID"`
	CampaignImageURL string    `json:"campaign_image_url"`
	ExpiredAt        time.Time `json:"expired_at"`
}
