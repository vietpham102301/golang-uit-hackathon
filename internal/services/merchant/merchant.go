package merchant

import (
	"errors"
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
	"gorm.io/gorm"
)

type IMerchant interface {
	ListMerchantByIDs(IDs []int64) ([]*models.Merchant, error)
	Create(record *models.Merchant) (*models.Merchant, error)
	GetByID(ID int) (*models.Merchant, error)
	ListByFilter(data map[string]interface{}) ([]*models.Merchant, error)
}

type Merchant struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewMerchant(appConfig *config.AppConfig, repo repos.IRepo) IMerchant {
	return &Merchant{
		cfg:  appConfig,
		repo: repo,
	}
}

func (m Merchant) ListMerchantByIDs(IDs []int64) ([]*models.Merchant, error) {
	records, err := m.repo.Merchant().ListMerchantByIDs(IDs)
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return nil, err
	}
	return records, nil
}

func (m Merchant) Create(record *models.Merchant) (*models.Merchant, error) {
	record, err := m.repo.Merchant().Create(record)
	return record, err
}

func (m *Merchant) GetByID(ID int) (*models.Merchant, error) {
	merchant, err := m.repo.Merchant().FindByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.Merchant{}, nil
		}

		return nil, err
	}
	return merchant, nil
}

func (m *Merchant) ListByFilter(data map[string]interface{}) ([]*models.Merchant, error) {
	listMerchant, err := m.repo.Merchant().ListByConditions(data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var merchants []*models.Merchant
			return merchants, nil
		}

		return nil, err
	}

	return listMerchant, nil
}
