package models

type Provider struct {
	ID           int64  `json:"id"`
	UserID       string `json:"user_id" gorm:"-"`
	ProviderName string `json:"provider_name"`
}
