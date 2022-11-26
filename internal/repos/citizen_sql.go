package repos

import (
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"gorm.io/gorm"
)

type CitizenSQLRepo struct {
	db *gorm.DB
}

func NewCitizenSQLRepo(db *gorm.DB) ICitizenRepo {
	return &CitizenSQLRepo{
		db: db,
	}
}

func (c CitizenSQLRepo) Create(record *models.Citizen) (*models.Citizen, error) {
	err := c.db.Create(record).Error
	return record, err
}

func (c CitizenSQLRepo) GetByID(ID int64) (*models.Citizen, error) {
	record := &models.Citizen{}
	err := c.db.Where("id = ?", ID).First(record).Error

	return record, err
}
