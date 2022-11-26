package models

type CreateUserRequest struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Role     int     `json:"role"`
	Citizen  Citizen `json:"citizen"`
}

type Citizen struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
