package models

import "github.com/vietpham1023/golang-uit-hackathon/internal/models"

type CreateUserResponse struct {
	ID              int64           `json:"id"`
	Email           string          `json:"email"`
	Role            int             `json:"role"`
	Token           string          `json:"token"`
	CitizenResponse CitizenResponse `json:"citizen"`
}

func (c *CreateUserResponse) ToUserResponse(user *models.User) *CreateUserResponse {
	return &CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Token: user.Token,
		CitizenResponse: CitizenResponse{
			ID:   user.Citizen.ID,
			Name: user.Citizen.Name,
		},
	}
}
