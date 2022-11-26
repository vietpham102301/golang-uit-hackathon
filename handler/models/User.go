package models

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	CitizenID int    `json:"citizen_id"`
}
