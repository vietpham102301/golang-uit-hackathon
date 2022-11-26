package models

type Campaign struct {
	ID         int64  `json:"id"`
	ProviderID int64  `json:"provider_id"`
	Name       string `json:"name"`
	Dsc        string `json:"dsc"` // description
	RuleID     int64  `json:"RuleID"`
}
