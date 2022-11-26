package merchant

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IMerchant interface {
	ListMerchantByIDs(IDs []int64) ([]*models.Merchant, error)
	Create(record *models.Merchant) (*models.Merchant, error)
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
