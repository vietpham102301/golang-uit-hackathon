package models

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Role      int      `json:"role"`
	Email     string   `json:"email"`
	Token     string   `json:"token"`
	CitizenID int64    `json:"citizen_id"`
	Citizen   *Citizen `json:"citizen" gorm:"-"`
}
