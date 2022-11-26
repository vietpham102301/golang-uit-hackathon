package book

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IBook interface {
	CreateBook(record *models.Book) (*models.Book, error)
	GetByID(ID int) (*models.Book, error)
	ListByFilter(data map[string]string) ([]*models.Book, error)
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

func (b *Book) GetByID(ID int) (*models.Book, error) {
	record, err := b.repo.Book().GetByID(ID)
	return record, err
}

func (b *Book) ListByFilter(data map[string]string) ([]*models.Book, error) {
	page, size, filter := getFilterList(data)
	records, err := b.repo.Book().ListBookByFilter(page, size, filter)
	if err != nil {
		fmt.Printf("list book fail with err: %v\n", err)
		return nil, err
	}
	return records, nil
}
