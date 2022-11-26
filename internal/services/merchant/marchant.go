package merchantService

import (
	"context"
	"errors"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/handler/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
	"gorm.io/gorm"
)

type IMerchant interface {
	CreateMerchant(ctx context.Context, record *models.Merchant) error
	GetByID(ctx context.Context, ID int) (*models.Merchant, error)
	ListByFilter(ctx context.Context, data map[string]interface{}) ([]*models.Merchant, error)
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

func (m *Merchant) CreateMerchant(ctx context.Context, record *models.Merchant) error {
	err := m.repo.Merchant().CreateMerchant(ctx, record)
	return err
}

func (m *Merchant) GetByID(ctx context.Context, ID int) (*models.Merchant, error) {
	merchant, err := m.repo.Merchant().FindByID(ctx, ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.Merchant{}, nil
		}

		return nil, err
	}
	return merchant, nil
}

func (m *Merchant) ListByFilter(ctx context.Context, data map[string]interface{}) ([]*models.Merchant, error) {
	listMerchant, err := m.repo.Merchant().ListByConditions(ctx, data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var merchants []*models.Merchant
			return merchants, nil
		}

		return nil, err
	}

	return listMerchant, nil
}
