package models

type Item struct {
	ID         int64  `json:"id"`
	ItemName   string `json:"item_name"`
	ProviderID string `json:"provider_id"`
}
