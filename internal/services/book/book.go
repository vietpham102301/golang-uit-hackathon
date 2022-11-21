package book

import (
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IBook interface {
	CreateBook(record *models.Book) (*models.Book, error)
}

type Book struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewBook(appConfig *config.AppConfig, repo repos.IRepo) IBook {
	return &Book{
		cfg:  appConfig,
		repo: repo,
	}
}

func (b *Book) CreateBook(record *models.Book) (*models.Book, error) {
	record, err := b.repo.Book().Create(record)
	return record, err
}
