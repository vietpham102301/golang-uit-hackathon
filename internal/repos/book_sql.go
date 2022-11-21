package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type BookSQLRepo struct {
	db *gorm.DB
}

func NewBookSQLRepo(db *gorm.DB) IBookRepo {
	return &BookSQLRepo{
		db: db,
	}
}

func (b *BookSQLRepo) Create(record *models.Book) (*models.Book, error) {
	err := b.db.Create(record).Error
	return record, err
}
