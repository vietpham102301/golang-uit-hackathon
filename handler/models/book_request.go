package models

type CreateBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
