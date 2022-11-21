package models

import "github.com/vietpham1023/golang-uit-hackathon/internal/models"

type CreateBookResponse struct {
	ID     int64
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (c *CreateBookResponse) ToBookResponse(book *models.Book) *CreateBookResponse {
	return &CreateBookResponse{
		ID:     book.ID,
		Name:   book.Name,
		Author: book.Author,
	}
}
