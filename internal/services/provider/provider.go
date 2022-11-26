package provider

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IProvider interface {
	ListProviderByIDs(IDs []int64) ([]*models.Provider, error)
	Create(record *models.Provider) (*models.Provider, error)
}

type Provider struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewProvider(appConfig *config.AppConfig, repo repos.IRepo) IProvider {
	return &Provider{
		cfg:  appConfig,
		repo: repo,
	}
}

func (r Provider) ListProviderByIDs(IDs []int64) ([]*models.Provider, error) {
	records, err := r.repo.Provider().ListProviderByIDs(IDs)
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return nil, err
	}
	return records, nil
}

func (r Provider) Create(record *models.Provider) (*models.Provider, error) {
	record, err := r.repo.Provider().Create(record)
	return record, err
}
