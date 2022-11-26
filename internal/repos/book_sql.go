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

func (b *BookSQLRepo) GetByID(ID int) (*models.Book, error) {
	record := &models.Book{}
	err := b.db.Where("id = ?", ID).First(record).Error

	return record, err
}

func (b *BookSQLRepo) ListBookByFilter(size int, page int, filter map[string]interface{}) ([]*models.Book, error) {
	var records []*models.Book
	query := b.db
	authorName, ok := filter["author"]
	if ok {
		query = query.Where("author = ?", authorName)
	}
	var err error
	if page != -1 && size != -1 {
		offset := (page - 1) * size
		err = query.Order("created_at DESC").Limit(size).Offset(offset).Find(&records).Error
	} else {
		err = query.Order("created_at DESC").Find(&records).Error
	}
	return records, err
}
