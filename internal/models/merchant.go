package models

type Merchant struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id" gorm:"-"`
	MerchantName string `json:"merchant_name"`
}
