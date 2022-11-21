package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
)

type IRepo interface {
	Book() IBookRepo
}

type IBookRepo interface {
	Create(record *models.Book) (*models.Book, error)
}
