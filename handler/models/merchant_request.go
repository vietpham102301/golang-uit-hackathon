package models

type Merchant struct {
	ID     int    `json:"id" gorm:"id"`
	UserID int    `json:"user_id" gorm:"user_id"`
	Name   string `json:"name" gorm:"name"`
}

type ListMerchant struct {
	ID   int    `form:"id"`
	User int    `form:"user_id"`
	Name string `form:"name"`
}

type GetMerchant struct {
	ID int `uri:"id" gorm:"id"`
}
